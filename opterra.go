package main

/*

infra.tf
payload.json (json description of vm)
resource.json (status tracking)
terraform.tfstate ()

*/

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"io/ioutil"

	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jeffail/gabs"
	"github.com/perigee/terrant/app"
)

const (
	terraBUCKETNAME    = "autoterrarepostate"
	terraDEFAULTREGION = "us-east-1"
	terraTerraformFile = "infra.tf"
	terraStatusFile    = "resource.json"
	terraAttributeFile = "attributes.json"
	terraSPACEIDKEY    = "ubispaceid"
)

func downloadS3object(s3client *s3.S3, key string) ([]byte, error) {
	res, err := s3client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(terraBUCKETNAME),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, res.Body); err != nil {
		return nil, err

	}

	return buf.Bytes(), nil
}

func uploadS3object(s3client *s3.S3, key string, body []byte) error {

	_, err := s3client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(terraBUCKETNAME),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func s3KeyGen(ctx *app.CreateChefContext, filename string) string {
	okey := make([]string, 3)

	okey[0] = ctx.Request.Header.Get(terraSPACEIDKEY)
	okey[1] = ctx.Payload.Vmuid
	okey[2] = filename
	return strings.Join(okey, "/")
}

// CreateS3Client creates the client
func CreateS3Client() *s3.S3 {
	sess, err := session.NewSession()

	if err != nil {
		panic(err)
	}

	return s3.New(sess, &aws.Config{Region: aws.String(terraDEFAULTREGION)})

}

// ChefCreateImp fetch the object
func ChefCreateImp(ctx *app.CreateChefContext) ([]byte, error) {

	svc := CreateS3Client()

	resp, err := UpdateTerraFile(ctx, svc)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return []byte(resp), nil
}

// CheckJobStatus verifies the status of current resource
func CheckJobStatus(ctx *app.CreateChefContext, s3client *s3.S3, key string) ([]byte, error) {

	res, err := downloadS3object(s3client, key)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	jsonObj, err := gabs.ParseJSON(res)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	value, ok := jsonObj.Path("status").Data().(string)

	if ok {
		return []byte(value), nil
	}

	return []byte(""), nil
}

// UpdateTerraFile updates
func UpdateTerraFile(ctx *app.CreateChefContext, s3client *s3.S3) ([]byte, error) {

	runlist := ctx.Payload.Runlist

	res, err := downloadS3object(s3client, s3KeyGen(ctx, terraTerraformFile))

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	jsonObj, err := gabs.ParseJSON(res)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	children, _ := jsonObj.S("module").Children()

	for _, child := range children {
		child.SetP(time.Now().UTC(), "single_vm.vm_trigger_hash")
		child.Array("single_vm", "vm_bootstrap_runlist")

		for _, runbook := range runlist {
			child.ArrayAppend(runbook, "single_vm", "vm_bootstrap_runlist")
		}
	}

	jsonStr := jsonObj.String()

	return []byte(jsonStr), nil
}

// UpdateContainerID updates the container id
func UpdateContainerID(ctx *app.CreateChefContext, s3client *s3.S3, id string) error {

	res, err := downloadS3object(s3client, s3KeyGen(ctx, terraStatusFile))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	jsonObj, err := gabs.ParseJSON(res)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	jsonObj.SetP(id, "internal.containerId")
	jsonStr := jsonObj.String()

	if err := uploadS3object(s3client, s3KeyGen(ctx, "tmp_resource.json"), []byte(jsonStr)); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

// JSON2File creates the file based on given json
/*
    if err := JSON2File(ctx.Payload.NodeAttributes, "atrb.json"); err != nil {
		panic(err)
	}
*/
func JSON2File(in *interface{}, filename string) error {
	jsonStr, err := json.Marshal(*in)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	ioutil.WriteFile(filename, jsonStr, 0644)
	return nil
}

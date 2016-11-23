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

	"io/ioutil"

	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/perigee/terrant/app"
	"github.com/perigee/terrant/design"
)

const (
	terraBUCKETNAME     = "autoterrarepostate"
	terraDEFAULTREGION  = "us-east-1"
	terraTerraformState = "infra.tf"
	terraStatusFile     = "resource.json"
	terraAttributeFile  = "attributes.json"
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

	okey[0] = ctx.Request.Header.Get(design.SPACEID)
	okey[1] = ctx.Payload.Vmuid
	okey[2] = filename
	return strings.Join(okey, "/")
}

// FetchObject fetch the object
func FetchObject(ctx *app.CreateChefContext) ([]byte, error) {
	sess, err := session.NewSession()

	if err != nil {
		panic(err)
	}

	svc := s3.New(sess, &aws.Config{Region: aws.String(terraDEFAULTREGION)})

	resp, err := downloadS3object(svc, s3KeyGen(ctx, terraStatusFile))

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if err := uploadS3object(svc, s3KeyGen(ctx, "jun_tmp_file"), resp); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return resp, nil
}

// CheckJobStatus verifies the status of current resource
func CheckJobStatus(ctx *app.CreateChefContext) (string, error) {

	return "", nil
}

// UpdateTerraFile updates
func UpdateTerraFile(s3obj *s3.GetObjectOutput) (*s3.GetObjectOutput, error) {
	return nil, nil
}

// UploadObject upload the object on s3
func UploadObject() error {
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

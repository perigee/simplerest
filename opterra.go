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
	BUCKETNAME     = "autoterrarepostate"
	DEFAULTREGION  = "us-east-1"
	TerraformState = "infra.tf"
)

func downloadS3object(s3client *s3.S3, key string) ([]byte, error) {
	res, err := s3client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("autoterrarepostate"),
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

// FetchObject fetch the object
func FetchObject(ctx *app.CreateChefContext) ([]byte, error) {

	//
	okey := make([]string, 3)

	okey[0] = ctx.Request.Header.Get(design.SPACEID)
	okey[1] = ctx.Payload.Vmuid
	okey[2] = "infra.tf"

	objectkey := strings.Join(okey, "/")

	sess, err := session.NewSession()

	if err != nil {
		panic(err)
	}

	svc := s3.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	resp, err := downloadS3object(svc, objectkey)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(resp)

	return resp, nil
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

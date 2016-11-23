package main

import (
	"encoding/json"
	"fmt"

	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	BUCKETNAME     = "autoterrarepostate"
	DEFAULTREGION  = "us-east-1"
	TerraformState = "infra.tf"
)

// FetchObject fetch the object
func FetchObject(spaceid string) (*s3.GetObjectOutput, error) {
	sess, err := session.NewSession()

	if err != nil {
		panic(err)
	}

	svc := s3.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	params := &s3.GetObjectInput{
		Bucket: aws.String("autoterrarepostate"),
		Key:    aws.String(spaceid),
	}

	resp, err := svc.GetObject(params)

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
func JSON2File(in *interface{}, filename string) error {
	jsonStr, err := json.Marshal(*in)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	ioutil.WriteFile(filename, jsonStr, 0644)
	return nil
}

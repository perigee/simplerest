package main

import (
	"io/ioutil"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

func GetVaraiblesList(file *ast.File) ([]string, error) {
	a := []string{"sef", "sdfe"}

	return a, nil
}

func HCLmapper(filename string) (*ast.File, error) {

	byteArray, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	hclAstFile, err := hcl.ParseBytes(byteArray)

	if err != nil {
		return nil, err
	}

	return hclAstFile, nil
}

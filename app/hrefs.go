//************************************************************************//
// API "provisioner": Application Resource Href Factories
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/perigee/terrant/design
// --out=$(GOPATH)/src/github.com/perigee/terrant
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"fmt"
	"strings"
)

// ChefHref returns the resource href.
func ChefHref(vmuid interface{}) string {
	paramvmuid := strings.TrimLeftFunc(fmt.Sprintf("%v", vmuid), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/provisioner/chef/%v", paramvmuid)
}

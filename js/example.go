//************************************************************************//
// provisioner JavaScript Client Example
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/perigee/terrant/design
// --out=$(GOPATH)/src/github.com/perigee/terrant
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package js

import "github.com/goadesign/goa"

// MountController mounts the JavaScript example controller under "/js".
// This is just an example, not the best way to do this. A better way would be to specify a file
// server using the Files DSL in the design.
// Use --noexample to prevent this file from being generated.
func MountController(service *goa.Service) {
	// Serve static files under js
	service.ServeFiles("/js/*filepath", "/home/go/src/github.com/perigee/terrant/js")
	service.LogInfo("mount", "ctrl", "JS", "action", "ServeFiles", "route", "GET /js/*")
}

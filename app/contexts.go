//************************************************************************//
// API "provisioner": Application Contexts
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
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// CreateChefContext provides the chef create action context.
type CreateChefContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *ChefPayload
}

// NewCreateChefContext parses the incoming request URL and body, performs validations and creates the
// context used by the chef controller create action.
func NewCreateChefContext(ctx context.Context, service *goa.Service) (*CreateChefContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateChefContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateChefContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "(^/[0-9]+")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

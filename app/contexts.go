//************************************************************************//
// API "infra": Application Contexts
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

// CreateResourceContext provides the resource create action context.
type CreateResourceContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ResourceID string
}

// NewCreateResourceContext parses the incoming request URL and body, performs validations and creates the
// context used by the resource controller create action.
func NewCreateResourceContext(ctx context.Context, service *goa.Service) (*CreateResourceContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateResourceContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramResourceID := req.Params["resourceID"]
	if len(paramResourceID) > 0 {
		rawResourceID := paramResourceID[0]
		rctx.ResourceID = rawResourceID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 201.
func (ctx *CreateResourceContext) OK() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

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
	Payload *CreateChefPayload
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

// createChefPayload is the chef create action payload.
type createChefPayload struct {
	// {docker: {name: dockername}}
	NodeAttributes *interface{} `form:"nodeAttributes,omitempty" json:"nodeAttributes,omitempty" xml:"nodeAttributes,omitempty"`
	Runlist        []string     `form:"runlist,omitempty" json:"runlist,omitempty" xml:"runlist,omitempty"`
	// kdielsie
	Vmuid *string `form:"vmuid,omitempty" json:"vmuid,omitempty" xml:"vmuid,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createChefPayload) Validate() (err error) {
	if payload.Vmuid == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "vmuid"))
	}
	if payload.Runlist == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "runlist"))
	}

	return
}

// Publicize creates CreateChefPayload from createChefPayload
func (payload *createChefPayload) Publicize() *CreateChefPayload {
	var pub CreateChefPayload
	if payload.NodeAttributes != nil {
		pub.NodeAttributes = payload.NodeAttributes
	}
	if payload.Runlist != nil {
		pub.Runlist = payload.Runlist
	}
	if payload.Vmuid != nil {
		pub.Vmuid = *payload.Vmuid
	}
	return &pub
}

// CreateChefPayload is the chef create action payload.
type CreateChefPayload struct {
	// {docker: {name: dockername}}
	NodeAttributes *interface{} `form:"nodeAttributes,omitempty" json:"nodeAttributes,omitempty" xml:"nodeAttributes,omitempty"`
	Runlist        []string     `form:"runlist" json:"runlist" xml:"runlist"`
	// kdielsie
	Vmuid string `form:"vmuid" json:"vmuid" xml:"vmuid"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateChefPayload) Validate() (err error) {
	if payload.Vmuid == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "vmuid"))
	}
	if payload.Runlist == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "runlist"))
	}

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateChefContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "(^/[0-9]+")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// ShowChefContext provides the chef show action context.
type ShowChefContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	VMUID *string
	Vmuid string
}

// NewShowChefContext parses the incoming request URL and body, performs validations and creates the
// context used by the chef controller show action.
func NewShowChefContext(ctx context.Context, service *goa.Service) (*ShowChefContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowChefContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramVMUID := req.Params["VMUID"]
	if len(paramVMUID) > 0 {
		rawVMUID := paramVMUID[0]
		rctx.VMUID = &rawVMUID
	}
	paramVmuid := req.Params["vmuid"]
	if len(paramVmuid) > 0 {
		rawVmuid := paramVmuid[0]
		rctx.Vmuid = rawVmuid
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowChefContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowChefContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

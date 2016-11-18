//************************************************************************//
// API "infra": Application Controllers
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
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ResourceController is the controller interface for the Resource actions.
type ResourceController interface {
	goa.Muxer
	Create(*CreateResourceContext) error
}

// MountResourceController "mounts" a Resource resource controller on the given service.
func MountResourceController(service *goa.Service, ctrl ResourceController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateResourceContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/infra/resource/:resourceID", ctrl.MuxHandler("Create", h, nil))
	service.LogInfo("mount", "ctrl", "Resource", "action", "Create", "route", "POST /infra/resource/:resourceID")
}

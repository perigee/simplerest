package main

import (
	"github.com/goadesign/goa"
	"github.com/perigee/terrant/app"
)

// ResourceController implements the resource resource.
type ResourceController struct {
	*goa.Controller
}

// NewResourceController creates a resource controller.
func NewResourceController(service *goa.Service) *ResourceController {
	return &ResourceController{Controller: service.NewController("ResourceController")}
}

// Create runs the create action.
func (c *ResourceController) Create(ctx *app.CreateResourceContext) error {
	// ResourceController_Create: start_implement

	// Put your logic here

	// ResourceController_Create: end_implement
	return nil
}

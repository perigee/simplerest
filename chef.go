package main

import (
	"github.com/goadesign/goa"
	"github.com/perigee/terrant/app"
)

// ChefController implements the chef resource.
type ChefController struct {
	*goa.Controller
}

// NewChefController creates a chef controller.
func NewChefController(service *goa.Service) *ChefController {
	return &ChefController{Controller: service.NewController("ChefController")}
}

// Create runs the create action.
func (c *ChefController) Create(ctx *app.CreateChefContext) error {
	// ChefController_Create: start_implement

	// Put your logic here
	res, _ := FetchObject(ctx)

	return ctx.OK([]byte(res))

	// ChefController_Create: end_implement
	return nil
}

// Show runs the show action.
func (c *ChefController) Show(ctx *app.ShowChefContext) error {
	// ChefController_Show: start_implement

	// Put your logic here

	// ChefController_Show: end_implement
	return nil
}

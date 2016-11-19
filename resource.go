package main

import (
	"github.com/goadesign/goa"
	"github.com/perigee/terrant/app"
	"github.com/fsouza/go-dockerclient"
	
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
	id := ctx.ResourceID


	ch := make(chan string)
	endpoint := "unix:///var/run/docker.sock"

	client, err := docker.NewClient(endpoint)


	if err != nil {
	   panic(err)
	}

	
	go func(id string) {

	   //containers, err := client.ListContainers(docker.ListContainersOptions{All: false})


	   config := new(docker.Config)
	   config.Image = id
	   
	   container, err := client.CreateContainer(docker.CreateContainerOptions{Name:"haha", Config: config}) 
	   
	   if err != nil {
	      panic(err)
	   }

	   if err := client.StartContainer(container.ID, new(docker.HostConfig)); err  != nil {
	      panic(err)
	   }

	   ch <- container.ID

	}(id)

	return ctx.OK([]byte(<-ch))

	// ResourceController_Create: end_implement
	return nil
}

package main

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
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

func getContainerConfig(imageID string) *container.Config {
	return &container.Config{
		Image: imageID,
	}
}

func getHostConfig() *container.HostConfig {
	return &container.HostConfig{}
}

func getNetworkingConfig() *network.NetworkingConfig {
	return &network.NetworkingConfig{}
}

// Create runs the create action.
func (c *ResourceController) Create(ctx *app.CreateResourceContext) error {
	// ResourceController_Create: start_implement

	// Put your logic here
	id := ctx.ResourceID

	ch := make(chan string)

	client, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	go func(id string) error {

		_, err := client.ImagePull(ctx, "nginx", types.ImagePullOptions{All: false})

		if err != nil {
			if err == IsErrImageNotFound(err)
				
			panic(err)
		}

		defer client.Close()

		container, err := client.ContainerCreate(ctx, getContainerConfig(id), getHostConfig(), getNetworkingConfig(), "haha")

		if err != nil {
			panic(err)
		}

		ch <- container.ID

		// create the docker container

	}(id)

	return ctx.OK([]byte(<-ch))

	// ResourceController_Create: end_implement
	return nil
}

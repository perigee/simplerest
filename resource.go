package main

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/goadesign/goa"
	"github.com/perigee/terrant/app"
	"golang.org/x/net/context"
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

	client, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	defer client.Close()

	ch := make(chan string)

	go func(id string) {

		//if len(matches) == 0 || strings.HasSuffix(id, "latest") {

		fmt.Printf(client.ClientVersion())

		resp, err := client.ImagePull(context.Background(), id, types.ImagePullOptions{})

		defer resp.Close()

		if err != nil {

			panic(err)

		} else {
			ch <- "not found"
		}

		return

		if err != nil {
			fmt.Printf("========================== NO IMAGE in imagepull")
			panic(err)
		}

		container, err := client.ContainerCreate(ctx, getContainerConfig(id), getHostConfig(),
			getNetworkingConfig(), "")

		if err != nil {
			fmt.Printf("========================== NO IMAGE in container")
			panic(err)
		}

		if err := client.ContainerStart(ctx, container.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}

		ch <- container.ID

		// create the docker container

	}(id)

	return ctx.OK([]byte(<-ch))

	// ResourceController_Create: end_implement
	return nil
}

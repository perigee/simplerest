package main

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/perigee/terrant/app"
)

func StartTerra(ctx *app.CreateChefContext) {

	endpoint := "http://127.0.0.1"

	client, err := docker.NewClient(endpoint)

	if err != nil {
		panic(err)
	}

	go func(ctx *app.CreateChefContext) {

		//containers, err := client.ListContainers(docker.ListContainersOptions{All: false})

		config := new(docker.Config)
		config.Image = ctx.Payload.Vmuid

		container, err := client.CreateContainer(docker.CreateContainerOptions{Config: config})

		if err != nil {
			panic(err)
		}

		if err := client.StartContainer(container.ID, &docker.HostConfig{}); err != nil {
			panic(err)
		}

	}(ctx)

}

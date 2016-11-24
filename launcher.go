package main

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
	"github.com/perigee/terrant/app"
)

// StartTerra starts the terra container
func StartTerra(ctx *app.CreateChefContext) {

	// use it
	endpoint := "http://127.0.0.1:2375"

	client, err := docker.NewClient(endpoint)

	if err != nil {
		panic(err)
	}

	go func(ctx *app.CreateChefContext) {

		//containers, err := client.ListContainers(docker.ListContainersOptions{All: false})

		config := new(docker.Config)
		config.Image = ctx.Payload.Vmuid
		config.Image = "nginx:alpine"

		container, err := client.CreateContainer(docker.CreateContainerOptions{Config: config})

		if err != nil {
			panic(err)
		}

		if err := client.StartContainer(container.ID, &docker.HostConfig{}); err != nil {
			panic(err)
		}

		s3client := CreateS3Client()

		if err := UpdateContainerID(ctx, s3client, container.ID); err != nil {
			fmt.Println(err.Error())
		}

	}(ctx)

}

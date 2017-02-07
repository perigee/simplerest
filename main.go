package main

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {

	configDocker := &container.Config{
		Image: "nginx:alpine",
	}

	cli, err := client.NewEnvClient()

	if err != nil {
		fmt.Printf("Create: %s", err.Error())
		return
	}

	filter := filters.NewArgs()

	filter.Add("name", "my_testing")

	option := types.NetworkListOptions{
		Filters: filter,
	}

	ctx := context.Background()

	netsx, err := cli.NetworkList(ctx, option)

	if err != nil {
		fmt.Printf("Problem: %s", err.Error())
	}

	var networkID string
	for _, netrs := range netsx {
		fmt.Printf("network: %s:%s", netrs.Name, netrs.ID)
		networkID = netrs.ID
	}

	contr, err := cli.ContainerCreate(ctx, configDocker, nil, nil, "")

	if err != nil {
		fmt.Printf("Creation error: %s", err.Error())
	}

	if err := cli.NetworkConnect(ctx, networkID, contr.ID, nil); err != nil {
		fmt.Printf("Network error: %s", err.Error())
	}

	if err := cli.ContainerStart(ctx, contr.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Printf("Start error: %s", err.Error())
	}
}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// PayloadCollection a collection of Payload
type PayloadCollection struct {
	WindowsVersion string    `json:"version"`
	Token          string    `json:"token"`
	Payloads       []Payload `json:"data"`
}

// Payload terraform docker task
type Payload struct {
	// [redacted]
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

// Provision launches the terraform docker container
func (p *Payload) Provision(id string) error {
	// Docker launch
	client, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	defer client.Close()

	fmt.Printf(client.ClientVersion())

	//resp, err := client.ImagePull(context.Background(), id, types.ImagePullOptions{})
	//defer resp.Close()
	//if err != nil {
	//	panic(err)
	//}

	container, err := client.ContainerCreate(context.Background(),
		getContainerConfig(id), getHostConfig(), getNetworkingConfig(), "")

	if err != nil {
		panic(err)
	}

	if err := client.ContainerStart(context.Background(), container.ID,
		types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	// tracking the container status
	return nil
}

var (
	MaxWorker = os.Getenv("MAX_WORKER")
	MaxQueue  = os.Getenv("MAX_Queue")
)

var JobQueue chan Job

type Job struct {
	Payload Payload
}

// Worker the worker launch the terrform container
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

// NewWorker returns a new worker
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start starts the worker
func (w Worker) Start(id string) {

	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				if err := job.Payload.Provision(id); err != nil {
					// Missing: Error handling
				}

			case <-w.quit:
				return

			}
		}
	}()

}

// Stop stops the worker
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()

}

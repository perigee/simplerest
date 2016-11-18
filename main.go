//go:generate goagen bootstrap -d github.com/perigee/terrant/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/perigee/terrant/app"
)

func main() {
	// Create service
	service := goa.New("infra")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "resource" controller
	c := NewResourceController(service)
	app.MountResourceController(service, c)

	// Start service
	if err := service.ListenAndServe(":8090"); err != nil {
		service.LogError("startup", "err", err)
	}
}

package design

import (
        . "github.com/goadesign/goa/design"
        . "github.com/goadesign/goa/design/apidsl"
)

var _ = API("infra", func() {
        Title("The adder API")
        Description("A teaser for goa")
        Host("localhost:8090")
        Scheme("http")
	BasePath("/infra")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("resource", func() {
	Action("create", func() {
		Routing(POST("/resource/:resourceID"))
		Description("create resource by its ID")
		Params(func() {
			Param("resourceID", String, "Resource ID")
		})
		Response("OK", func() {
			Description("Return in case of creation")
			Status(201)
		})
	})
})

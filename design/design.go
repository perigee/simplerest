package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// ChefPayload constructs request body
var ChefPayload = Type("ChefPayload", func() {
	Attribute("vm_uid", String, "kdielsie")
	Attribute("nodeAttributes", String, "{docker: {name: dockername}}")
	Attribute("runlist", ArrayOf(String))
})

var _ = API("provisioner", func() {
	Title("The adder API")
	Description("A teaser for goa")
	Host("localhost:3001")
	Scheme("http")
	BasePath("/provisioner")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("chef", func() {
	Action("create", func() {
		Routing(POST(""))
		Payload(ChefPayload, func() {
			Required("vm_uid", "runlist")
		})
		Response(OK, "(^/[0-9]+")
	})

	Action("show", func() {
		Routing(
			GET("/:vm_uid"),
		)
		Description("Retrieve the status by VM ID")
		Params(func() {
			Param("vmID", String, "VM ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
	})
})

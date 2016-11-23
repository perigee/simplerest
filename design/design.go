package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

const SPACEID = "ubispaceid"

// ChefPayload constructs request body
var ChefPayload = Type("ChefPayload", func() {
	Attribute("vmuid", String, "kdielsie")
	Attribute("nodeAttributes", String, "{docker: {name: dockername}}")
	Attribute("runlist", ArrayOf(String))
})

var _ = API("provisioner", func() {
	Title("Provisioner API")
	Description("A provisioner for VM")
	Host("localhost:3001")
	Scheme("http")
	BasePath("/provisioner")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("chef", func() {
	BasePath("/chef")
	Action("create", func() {
		Routing(POST(""))
		Payload(ChefPayload, func() {
			Required("vmuid", "runlist")
		})
		Response(OK, "(^/[0-9]+")
	})

	Action("show", func() {
		Routing(
			GET("/:vmuid"),
		)
		Description("Retrieve the status by VM ID")
		Params(func() {
			Param("VMUID", String, "VM ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

package design

import (
       . "github.com/goadesign/goa/design"
       . "github.com/goadesign/goa/design/apidsl"
)

var _ = API('config', func() {
    Routing(POST)
    
}) 
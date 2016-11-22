//************************************************************************//
// API "provisioner": Application User Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/perigee/terrant/design
// --out=$(GOPATH)/src/github.com/perigee/terrant
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

// chefPayload user type.
type chefPayload struct {
	// {docker: {name: dockername}}
	NodeAttributes *string  `form:"nodeAttributes,omitempty" json:"nodeAttributes,omitempty" xml:"nodeAttributes,omitempty"`
	Runlist        []string `form:"runlist,omitempty" json:"runlist,omitempty" xml:"runlist,omitempty"`
	// kdielsie
	UID *string `form:"uid,omitempty" json:"uid,omitempty" xml:"uid,omitempty"`
}

// Publicize creates ChefPayload from chefPayload
func (ut *chefPayload) Publicize() *ChefPayload {
	var pub ChefPayload
	if ut.NodeAttributes != nil {
		pub.NodeAttributes = ut.NodeAttributes
	}
	if ut.Runlist != nil {
		pub.Runlist = ut.Runlist
	}
	if ut.UID != nil {
		pub.UID = ut.UID
	}
	return &pub
}

// ChefPayload user type.
type ChefPayload struct {
	// {docker: {name: dockername}}
	NodeAttributes *string  `form:"nodeAttributes,omitempty" json:"nodeAttributes,omitempty" xml:"nodeAttributes,omitempty"`
	Runlist        []string `form:"runlist,omitempty" json:"runlist,omitempty" xml:"runlist,omitempty"`
	// kdielsie
	UID *string `form:"uid,omitempty" json:"uid,omitempty" xml:"uid,omitempty"`
}

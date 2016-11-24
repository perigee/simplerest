package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateChefPayload is the chef create action payload.
type CreateChefPayload struct {
	// {docker: {name: dockername}}
	NodeAttributes *interface{} `form:"nodeAttributes,omitempty" json:"nodeAttributes,omitempty" xml:"nodeAttributes,omitempty"`
	Runlist        []string     `form:"runlist" json:"runlist" xml:"runlist"`
	// kdielsie
	Vmuid string `form:"vmuid" json:"vmuid" xml:"vmuid"`
}

// CreateChefPath computes a request path to the create action of chef.
func CreateChefPath() string {
	return fmt.Sprintf("/api/v1/provisioner/chef")
}

// CreateChef makes a request to the create action endpoint of the chef resource
func (c *Client) CreateChef(ctx context.Context, path string, payload *CreateChefPayload) (*http.Response, error) {
	req, err := c.NewCreateChefRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateChefRequest create the request corresponding to the create action endpoint of the chef resource.
func (c *Client) NewCreateChefRequest(ctx context.Context, path string, payload *CreateChefPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowChefPath computes a request path to the show action of chef.
func ShowChefPath(vmuid string) string {
	return fmt.Sprintf("/api/v1/provisioner/chef/%v", vmuid)
}

// Retrieve the status by VM ID
func (c *Client) ShowChef(ctx context.Context, path string, vMUID *string) (*http.Response, error) {
	req, err := c.NewShowChefRequest(ctx, path, vMUID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowChefRequest create the request corresponding to the show action endpoint of the chef resource.
func (c *Client) NewShowChefRequest(ctx context.Context, path string, vMUID *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if vMUID != nil {
		values.Set("VMUID", *vMUID)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

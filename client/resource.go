package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateResourcePath computes a request path to the create action of resource.
func CreateResourcePath(resourceID string) string {
	return fmt.Sprintf("/infra/resource/%v", resourceID)
}

// create resource by its ID
func (c *Client) CreateResource(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewCreateResourceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateResourceRequest create the request corresponding to the create action endpoint of the resource resource.
func (c *Client) NewCreateResourceRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

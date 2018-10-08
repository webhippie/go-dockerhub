package dockerhub

import (
	"net/http"
)

// Client is a client for the DockerHub API.
type Client struct {
	httpClient *http.Client
	username   string
	password   string

	Organization OrganizationClient
	Repository   RepositoryClient
}

// A ClientOption is used to configure a Client.
type ClientOption func(*Client)

// WithUsername configures a Client to use the specified username for authentication.
func WithUsername(username string) ClientOption {
	return func(client *Client) {
		client.username = username
	}
}

// WithPassword configures a Client to use the specified password for authentication.
func WithPassword(password string) ClientOption {
	return func(client *Client) {
		client.password = password
	}
}

// NewClient creates a new client.
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		httpClient: &http.Client{},
	}

	for _, option := range options {
		option(client)
	}

	client.Organization = OrganizationClient{client: client}
	client.Repository = RepositoryClient{client: client}

	return client
}

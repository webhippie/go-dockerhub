package dockerhub

import (
	"context"
	"fmt"

	"github.com/webhippie/go-dockerhub/dockerhub/internal"
)

const (
	// NamespaceURL defines the namespace API endpoint.
	NamespaceURL = "https://hub.docker.com/v2/repositories/%s/"
)

// NamespaceClient is a client for the namespaces API.
type NamespaceClient struct {
	client *Client
}

// Get returns a specific namespace.
func (c *NamespaceClient) Get(ctx context.Context, name string) (*Namespace, error) {
	record := &Namespace{
		Name: name,
	}

	path := fmt.Sprintf(NamespaceURL, name)

	for {
		req, err := c.client.NewRequest(ctx, "GET", path, nil)

		if err != nil {
			return nil, err
		}

		result := internal.Namespace{}

		if _, err := c.client.Do(req, &result); err != nil {
			return nil, err
		}

		record.Count = result.Count

		for _, repository := range result.Repositories {
			record.Repositories = append(
				record.Repositories,
				RepositoriyFromInternal(repository),
			)
		}

		if result.Next == "" {
			break
		}

		path = result.Next
	}

	return record, nil
}

// Namespace represents an namespace to outside.
type Namespace struct {
	Name         string        `json:"name"`
	Count        int           `json:"count"`
	Repositories []*Repository `json:"repositories"`
}

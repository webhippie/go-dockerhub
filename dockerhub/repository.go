package dockerhub

import (
	"context"
	"fmt"
	"time"

	"github.com/webhippie/go-dockerhub/dockerhub/internal"
)

const (
	// RepositoryURL defines the repository API endpoint.
	RepositoryURL = "https://hub.docker.com/v2/repositories/%s/%s/"
)

// RepositoryClient is a client for the repositories API.
type RepositoryClient struct {
	client *Client
}

// Get returns a specific repository.
func (c *RepositoryClient) Get(ctx context.Context, namespace, repository string) (*Repository, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf(RepositoryURL, namespace, repository), nil)

	if err != nil {
		return nil, err
	}

	result := internal.Repository{}

	if _, err := c.client.Do(req, &result); err != nil {
		return nil, err
	}

	return RepositoriyFromInternal(result), nil
}

// Repository represents a repository to outside.
type Repository struct {
	Namespace   string    `json:"namespace"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
	Private     bool      `json:"private"`
	Automated   bool      `json:"automated"`
	Editable    bool      `json:"editable"`
	Migrated    bool      `json:"migrated"`
	Starred     bool      `json:"starred"`
	Stars       int       `json:"stars"`
	Pulls       int       `json:"pulls"`
	Updated     time.Time `json:"updated"`
}

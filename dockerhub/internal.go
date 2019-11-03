package dockerhub

import (
	"github.com/webhippie/go-dockerhub/dockerhub/internal"
)

// RepositoriyFromInternal converts an internal repository.
func RepositoriyFromInternal(r internal.Repository) *Repository {
	repository := &Repository{
		Namespace:   r.Namespace,
		Name:        r.Name,
		Type:        r.Type,
		Status:      r.Status,
		Description: r.Description,
		Private:     r.Private,
		Automated:   r.Automated,
		Editable:    r.Editable,
		Migrated:    r.Migrated,
		Starred:     r.Starred,
		Stars:       r.Stars,
		Pulls:       r.Pulls,
		Updated:     r.Updated,
	}

	return repository
}

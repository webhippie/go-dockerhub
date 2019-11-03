package internal

import (
	"time"
)

// Namespace represents an namespace within DockerHub.
type Namespace struct {
	Count        int          `json:"count"`
	Next         string       `json:"next"`
	Prev         string       `json:"previous"`
	Repositories []Repository `json:"results"`
}

// Repository represents a repository within DockerHub.
type Repository struct {
	User        string    `json:"user"`
	Namespace   string    `json:"namespace"`
	Name        string    `json:"name"`
	Type        string    `json:"repository_type"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
	Private     bool      `json:"is_private"`
	Automated   bool      `json:"is_automated"`
	Editable    bool      `json:"can_edit"`
	Migrated    bool      `json:"is_migrated"`
	Starred     bool      `json:"has_starred"`
	Stars       int       `json:"star_count"`
	Pulls       int       `json:"pull_count"`
	Updated     time.Time `json:"last_updated"`
}

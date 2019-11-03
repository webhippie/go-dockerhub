package main

import (
	"context"
	"log"

	"github.com/webhippie/go-dockerhub/dockerhub"
)

func main() {
	client := dockerhub.NewClient(
		dockerhub.WithUsername("octocat"),
		dockerhub.WithPassword("p455w0rd"),
	)

	repo, err := client.Repository.Get(
		context.Background(),
		"webhippie",
		"alpine",
	)

	if err != nil {
		log.Fatalf("Failed to get repository: %s", err)
	}

	if repo != nil {
		log.Printf("Found %s/%s image:", repo.Namespace, repo.Name)

		log.Printf("%d stars", repo.Stars)
		log.Printf("%d pulls", repo.Pulls)
	} else {
		log.Printf("Can't find repository webhippie/alpine")
	}
}

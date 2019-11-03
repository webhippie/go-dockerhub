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

	org, err := client.Namespace.Get(
		context.Background(),
		"webhippie",
	)

	if err != nil {
		log.Fatalf("Failed to get namespace: %s", err)
	}

	if org != nil {
		log.Printf("Found %d images for %s:", org.Count, org.Name)

		for _, repository := range org.Repositories {
			log.Printf("%s", repository.Name)
		}
	} else {
		log.Printf("Can't find namespace webhippie")
	}
}

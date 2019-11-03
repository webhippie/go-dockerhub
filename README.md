# go-dockerhub

[![Build Status](http://cloud.drone.io/api/badges/webhippie/go-dockerhub/status.svg)](http://cloud.drone.io/webhippie/go-dockerhub)
[![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/6c48092ea50d4c69a0a40253fba64f54)](https://www.codacy.com/manual/webhippie/go-dockerhub?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/go-dockerhub&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/webhippie/go-dockerhub?status.svg)](http://godoc.org/github.com/webhippie/go-dockerhub)
[![Go Report](http://goreportcard.com/badge/github.com/webhippie/go-dockerhub)](http://goreportcard.com/report/github.com/webhippie/go-dockerhub)

This package provides helpers related to DockerHub.

## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.12.


```bash
git clone https://github.com/webhippie/go-dockerhub.git
cd go-dockerhub

make clean generate test
```

## Examples

### Namespace

[embedmd]:# (examples/namespace/main.go go)
```go
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
```

### Repository

[embedmd]:# (examples/repository/main.go go)
```go
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
```

## Security

If you find a security issue please contact thomas@webhippie.de first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

*   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

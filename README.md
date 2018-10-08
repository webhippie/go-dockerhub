# Library for DockerHub

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/go-dockerhub/status.svg)](http://github.dronehippie.de/webhippie/go-dockerhub)
[![Stories in Ready](https://badge.waffle.io/webhippie/go-dockerhub.svg?label=ready&title=Ready)](http://waffle.io/webhippie/go-dockerhub)
[![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/6c48092ea50d4c69a0a40253fba64f54)](https://www.codacy.com/app/webhippie/go-dockerhub?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/go-dockerhub&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/webhippie/go-dockerhub?status.svg)](http://godoc.org/github.com/webhippie/go-dockerhub)
[![Go Report](http://goreportcard.com/badge/github.com/webhippie/go-dockerhub)](http://goreportcard.com/report/github.com/webhippie/go-dockerhub)

This repository will provides helpers related to DockerHub.

## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.8. It is also possible to just simply execute the `go get github.com/webhippie/go-dockerhub/...` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/webhippie/go-dockerhub/...
cd $GOPATH/src/github.com/webhippie/go-dockerhub
make retool sync clean generate test
```

## Examples

### Dummy

[embedmd]:# (examples/dummy/main.go go)
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Dummy")
}
```

## Security

If you find a security issue please contact thomas@webhippie.de first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

* [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

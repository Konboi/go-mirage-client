# go-mirage-client

[mirage](https://github.com/acidlemon/mirage) provide [API](https://github.com/acidlemon/mirage/blob/master/webapi.go#L24-L30)

this is a client library ot the mirage api

# Useage

`go get github.com/Konboi/go-mirage-client`


```go
package main

import (
	"fmt"
	"log"

	"github.com/Konboi/go-mirage-client"
)

const (
	MIRAGE_ENDPOINT = "http://mirage.example.com"
)

func main() {
	cli, err := mirage.NewClient(MIRAGE_ENDPOINT)
	if err != nil {
		log.Println(err.Error())
	}

	list, err := cli.List()

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("container num is:", len(list))
}
```

See `example/example.go`

# TODO

- command line tools

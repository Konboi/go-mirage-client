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
	cli := mirage.NewClinet(MIRAGE_ENDPOINT)

	list, err := cli.List()

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("container num is:", len(list.Result))

	params := make(map[string]string)
	params["branch"] = "somebranch"

	err = cli.Launch("subdomain", "image-name", params)
	if err != nil {
		log.Println(err.Error())
	}

	err = cli.Terminate("subdomain")
	if err != nil {
		log.Println(err.Error())
	}
}

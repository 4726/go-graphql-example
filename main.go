package main

import (
	"log"

	"github.com/4726/go-graphql-example/web"
)

func main() {
	log.Print("starting server")
	log.Print(web.RunServer())
}

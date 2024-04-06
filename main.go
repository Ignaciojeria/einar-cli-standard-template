package main

import (
	"archetype/app/infrastructure/server"
	"log"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

type Mock struct {
}

func main() {
	if err := ioc.LoadDependencies(); err != nil {
		log.Fatal(err)
	}
	server.Start()
}

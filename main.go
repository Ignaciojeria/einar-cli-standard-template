package main

import (
	_ "archetype/app/configuration"
	"archetype/app/infrastructure/server"
	"log"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

func main() {
	if err := ioc.LoadDependencies(); err != nil {
		log.Fatal(err)
	}
	ioc.Get[*server.Router](server.NewRouter).ServeHTTP()
}

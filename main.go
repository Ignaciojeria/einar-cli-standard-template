package main

import (
	"archetype/app/infrastructure/serverwrapper"
	"log"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

func main() {
	if err := ioc.LoadDependencies(); err != nil {
		log.Fatal(err)
	}
	serverwrapper.Start()
}

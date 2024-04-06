package main

import (
	"archetype/app/infrastructure/server"
	"log"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

func main() {
	ioc.Registry(echo.New)
	if err := ioc.LoadDependencies(); err != nil {
		log.Fatal(err)
	}
	server.Start()
}

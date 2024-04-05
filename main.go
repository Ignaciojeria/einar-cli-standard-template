package main

import (
	"archetype/app/configuration"
	"log"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

func main() {
	ioc.Registry(echo.New)
	if err := ioc.LoadDependencies(); err != nil {
		log.Fatal(err)
	}
	startServer()
}

func startServer() {
	e := ioc.Get[*echo.Echo](echo.New)

	routes := e.Routes()
	for _, route := range routes {
		log.Printf("Method: %s, Path: %s, Name: %s\n", route.Method, route.Path, route.Name)
	}

	port := ioc.Get[configuration.Conf](configuration.NewConf).PORT
	e.Logger.Fatal(e.Start(":" + port))
}

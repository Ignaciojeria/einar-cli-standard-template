package main

import (
	"archetype/app/configuration"
	_ "archetype/app/configuration"
	"log"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

func main() {
	ioc.Registry(echo.New)
	if err := ioc.LoadDependencies(); err != nil {
		log.Fatal(err)
	}
	e := ioc.Get[*echo.Echo](echo.New)
	port := ioc.Get[configuration.Conf](configuration.NewConf).PORT
	e.Logger.Fatal(e.Start(":" + port))
}

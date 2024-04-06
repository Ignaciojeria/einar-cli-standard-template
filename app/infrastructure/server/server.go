package server

import (
	"archetype/app/configuration"
	"log"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(echo.New)
	ioc.Registry(newServer, echo.New, configuration.NewConf)
}

func newServer(e *echo.Echo, conf configuration.Conf) server {
	return server{
		e:    e,
		conf: conf,
	}
}

type server struct {
	e    *echo.Echo
	conf configuration.Conf
}

func Start() {
	ioc.Get[server](newServer).start()
}

func (s server) start() {
	s.printRoutes()
	s.e.Logger.Fatal(s.e.Start(":" + s.conf.PORT))
}

func (s server) printRoutes() {
	routes := s.e.Routes()
	for _, route := range routes {
		log.Printf("Method: %s, Path: %s, Name: %s\n", route.Method, route.Path, route.Name)
	}
}

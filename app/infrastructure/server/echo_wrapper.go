package server

import (
	"archetype/app/configuration"
	"log"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(NewEchoWrapper, configuration.NewConf)
}

func NewEchoWrapper(conf configuration.Conf) EchoWrapper {
	return EchoWrapper{
		Echo: echo.New(),
		conf: conf,
	}
}

type EchoWrapper struct {
	*echo.Echo
	conf configuration.Conf
}

func Start() {
	ioc.Get[EchoWrapper](NewEchoWrapper).start()
}

func (s EchoWrapper) start() {
	s.printRoutes()
	s.Echo.Logger.Fatal(s.Echo.Start(":" + s.conf.PORT))
}

func (s EchoWrapper) printRoutes() {
	routes := s.Echo.Routes()
	for _, route := range routes {
		log.Printf("Method: %s, Path: %s, Name: %s\n", route.Method, route.Path, route.Name)
	}
}

package serverwrapper

import (
	"archetype/app/configuration"
	"log"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

type EchoWrapper struct {
	*echo.Echo
}

func init() {
	ioc.Registry(echo.New)
	ioc.Registry(NewEchoWrapper, echo.New)
}
func NewEchoWrapper(e *echo.Echo) EchoWrapper {
	return EchoWrapper{
		Echo: e,
	}
}

func Start() {
	ioc.Get[EchoWrapper](NewEchoWrapper).start()
}

func (s EchoWrapper) start() {
	s.printRoutes()
	s.Echo.Logger.Fatal(s.Echo.Start(":" + configuration.Values().PORT))
}

func (s EchoWrapper) printRoutes() {
	routes := s.Echo.Routes()
	for _, route := range routes {
		log.Printf("Method: %s, Path: %s, Name: %s\n", route.Method, route.Path, route.Name)
	}
}

package server

import (
	"archetype/app/configuration"
	"log"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

type EchoWrapper struct {
	*echo.Echo
	conf configuration.Conf
}

func init() {
	ioc.Registry(NewEchoWrapper, echo.New, configuration.NewConf)
}
func NewEchoWrapper(e *echo.Echo, conf configuration.Conf) EchoWrapper {
	return EchoWrapper{
		Echo: e,
		conf: conf,
	}
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

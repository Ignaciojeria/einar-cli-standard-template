package api

import (
	"archetype/app/infrastructure/server"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

type templateDelete struct {
}

func init() {
	ioc.Registry(newTemplateDelete, server.NewEchoWrapper)
}
func newTemplateDelete(e server.EchoWrapper) {
	e.DELETE("/insert-your-custom-pattern-here", templateDelete{}.handle)
}

func (api templateDelete) handle(c echo.Context) error {
	return c.String(http.StatusOK, "Unimplemented")
}

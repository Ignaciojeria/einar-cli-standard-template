package api

import (
	"archetype/app/infrastructure/server"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(newTemplateDelete, server.NewEchoWrapper)
}

type templateDelete struct {
}

func newTemplateDelete(e server.EchoWrapper) {
	e.DELETE("/insert-your-custom-pattern-here", templateDelete{}.handle)
}

func (api templateDelete) handle(c echo.Context) error {
	return c.String(http.StatusOK, "Unimplemented")
}

package api

import (
	"archetype/app/infrastructure/server"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(newTemplatePut, server.NewEchoWrapper)
}

type templatePut struct {
}

func newTemplatePut(e server.EchoWrapper) {
	e.PUT("/insert-your-custom-pattern-here", templatePut{}.handle)
}

func (api templatePut) handle(c echo.Context) error {
	return c.String(http.StatusOK, "Unimplemented")
}

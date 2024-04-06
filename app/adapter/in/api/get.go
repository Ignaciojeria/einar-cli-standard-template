package api

import (
	"archetype/app/infrastructure/server"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

type templateGet struct {
}

func init() {
	ioc.Registry(newTemplateGet, server.NewEchoWrapper)
}
func newTemplateGet(e server.EchoWrapper) {
	e.GET("/insert-your-custom-pattern-here", templateGet{}.handle)
}

func (api templateGet) handle(c echo.Context) error {
	return c.String(http.StatusOK, "Unimplemented")
}

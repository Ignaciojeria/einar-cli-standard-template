package api

import (
	"archetype/app/infrastructure/server"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

type templatePatch struct {
}

func init() {
	ioc.Registry(newTemplatePatch, server.NewEchoWrapper)
}
func newTemplatePatch(e server.EchoWrapper) {
	e.PATCH("/insert-your-custom-pattern-here", templatePatch{}.handle)
}

func (api templatePatch) handle(c echo.Context) error {
	return c.String(http.StatusOK, "Unimplemented")
}

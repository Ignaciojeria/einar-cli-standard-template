package api

import (
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(newTemplatePatch, echo.New)
}

type templatePatch struct {
}

func newTemplatePatch(e *echo.Echo) {
	e.PATCH("/insert-your-custom-pattern-here", templatePatch{}.handle)
}

func (api templatePatch) handle(c echo.Context) error {
	return c.String(http.StatusOK, "Unimplemented")
}

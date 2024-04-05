package api

import (
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(newTemplateGet, echo.New)
}

type templateGet struct {
}

func newTemplateGet(e *echo.Echo) {
	e.GET("/insert-your-custom-pattern-here", templateGet{}.handle)
}

func (api templateGet) handle(c echo.Context) error {
	return c.String(http.StatusOK, "Unimplemented")
}

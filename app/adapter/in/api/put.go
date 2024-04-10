package api

import (
	"archetype/app/infrastructure/serverwrapper"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(newTemplatePut, serverwrapper.NewEchoWrapper)
}
func newTemplatePut(e serverwrapper.EchoWrapper) {
	e.PUT("/insert-your-custom-pattern-here", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Unimplemented",
		})
	})
}

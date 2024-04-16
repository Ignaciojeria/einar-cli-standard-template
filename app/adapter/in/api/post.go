package api

import (
	"archetype/app/shared/infrastructure/serverwrapper"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(newTemplatePost, serverwrapper.NewEchoWrapper)
}
func newTemplatePost(e serverwrapper.EchoWrapper) {
	e.POST("/insert-your-custom-pattern-here", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Unimplemented",
		})
	})
}

package api

import (
	"archetype/app/infrastructure/server"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

func init() {
	ioc.Registry(newTemplatePost, server.NewEchoWrapper)
}

type templatePost struct {
}

func newTemplatePost(e server.EchoWrapper) {
	e.POST("/insert-your-custom-pattern-here", templatePost{}.handle)
}

func (api templatePost) handle(c echo.Context) error {
	return c.String(http.StatusOK, "Unimplemented")
}

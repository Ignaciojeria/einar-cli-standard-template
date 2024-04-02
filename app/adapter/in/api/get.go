package api

import (
	"archetype/app/infrastructure/server"
	"fmt"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

var _ = ioc.Registry(newTemplateGet, server.NewRouter)

type templateGet struct {
}

func newTemplateGet(router *server.Router) {
	router.GET("/insert-your-custom-pattern-here", templateGet{}.handle)
}

func (api templateGet) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unimplemented")
}

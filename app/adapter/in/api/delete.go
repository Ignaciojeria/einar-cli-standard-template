package api

import (
	"archetype/app/infrastructure/server"
	"fmt"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

var _ = ioc.Registry(newTemplateDelete, server.NewRouter)

type templateDelete struct {
}

func newTemplateDelete(router server.Router) {
	router.DELETE("/insert-your-custom-pattern-here", templateDelete{}.handle)
}

func (api templateDelete) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unimplemented")
}

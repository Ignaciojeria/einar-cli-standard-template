package api

import (
	"archetype/app/infrastructure/server"
	"fmt"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

var _ = ioc.Registry(newTemplatePut, server.NewRouter)

type templatePut struct {
}

func newTemplatePut(router server.Router) {
	router.PUT("/insert-your-custom-pattern-here", templatePut{}.handle)
}

func (api templatePut) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unimplemented")
}

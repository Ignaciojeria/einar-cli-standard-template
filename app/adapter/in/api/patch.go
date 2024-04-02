package api

import (
	"archetype/app/infrastructure/server"
	"fmt"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

var _ = ioc.Registry(newTemplatePatch, server.NewRouter)

type templatePatch struct {
}

func newTemplatePatch(router server.Router) {
	router.PATCH("/insert-your-custom-pattern-here", templatePatch{}.handle)
}

func (api templatePatch) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unimplemented")
}

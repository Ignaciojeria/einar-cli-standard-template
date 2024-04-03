package api

import (
	"archetype/app/infrastructure/server"
	"fmt"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

func init() {
	ioc.Registry(newTemplatePost, server.NewRouter)
}

type templatePost struct {
}

func newTemplatePost(router *server.Router) {
	router.POST("/insert-your-custom-pattern-here", templatePost{}.handle)
}

func (api templatePost) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unimplemented")
}

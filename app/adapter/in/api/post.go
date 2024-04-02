package api

import (
	"archetype/app/configuration"
	"fmt"
	"log"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

var _ = ioc.Registry(newTemplatePost, configuration.NewConf)

type templatePost struct {
}

func newTemplatePost(conf configuration.Conf) {
	adapter := templatePost{}
	pattern := http.MethodPost + " " + conf.ApiPrefix +
		"/insert-your-custom-pattern-here"
	log.Println(pattern)
	http.HandleFunc(pattern, adapter.handle)
}

func (api templatePost) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unimplemented")
}

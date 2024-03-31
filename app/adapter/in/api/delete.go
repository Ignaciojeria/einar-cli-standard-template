package api

import (
	"archetype/app/configuration"
	"fmt"
	"log"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

var _ = ioc.Registry(newTemplateDelete, configuration.NewConf)

type templateDelete struct {
}

func newTemplateDelete(conf configuration.Conf) error {
	adapter := templateDelete{}
	pattern := http.MethodDelete + " " + conf.ApiPrefix +
		"/insert-your-custom-pattern-here"
	log.Println(pattern)
	http.HandleFunc(pattern, adapter.handle)
	return nil
}

func (api templateDelete) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unimplemented")
}

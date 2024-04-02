package api

import (
	"archetype/app/configuration"
	"fmt"
	"log"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

var _ = ioc.Registry(newTemplatePut, configuration.NewConf)

type templatePut struct {
}

func newTemplatePut(conf configuration.Conf) {
	adapter := templatePut{}
	pattern := http.MethodPut + " " + conf.ApiPrefix +
		"/insert-your-custom-pattern-here"
	log.Println(pattern)
	http.HandleFunc(pattern, adapter.handle)
}

func (api templatePut) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unimplemented")
}

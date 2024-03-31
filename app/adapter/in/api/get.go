package api

import (
	"archetype/app/configuration"
	"fmt"
	"log"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

var _ = ioc.Registry(newTemplateGet, configuration.NewConf)

type templateGet struct {
}

func newTemplateGet(conf configuration.Conf) error {
	adapter := templateGet{}
	pattern := http.MethodGet + " " + conf.ApiPrefix +
		"/insert-your-custom-pattern-here"
	log.Println(pattern)
	http.HandleFunc(pattern, adapter.handle)
	return nil
}

func (api templateGet) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unimplemented")
}

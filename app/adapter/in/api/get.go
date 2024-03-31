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
	get := templateGet{}
	pattern := http.MethodGet + " " + conf.ApiPrefix +
		"/insert-your-custom-pattern-here"
	log.Println(pattern)
	http.HandleFunc(pattern, get.handler)
	return nil
}

func (api templateGet) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unimplemented")
}

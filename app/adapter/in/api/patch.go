package api

import (
	"archetype/app/configuration"
	"fmt"
	"log"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

var _ = ioc.Registry(newTemplatePatch, configuration.NewConf)

type templatePatch struct {
}

func newTemplatePatch(conf configuration.Conf) error {
	adapter := templatePatch{}
	pattern := http.MethodPatch + " " + conf.ApiPrefix +
		"/insert-your-custom-pattern-here"
	log.Println(pattern)
	http.HandleFunc(pattern, adapter.handle)
	return nil
}

func (api templatePatch) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unimplemented")
}

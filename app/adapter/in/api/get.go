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
	templateGet := templateGet{}
	path := http.MethodGet + " " + conf.ApiPrefix
	log.Println(path)
	http.HandleFunc(path, templateGet.handler)
	return nil
}

func (api templateGet) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola, este es un servidor HTTP básico en Go!")
}

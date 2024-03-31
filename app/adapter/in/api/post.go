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

func newTemplatePost(conf configuration.Conf) error {
	templatePost := templatePost{}
	path := http.MethodPost + " " + conf.ApiPrefix
	log.Println(path)
	http.HandleFunc(path, templatePost.handler)
	return nil
}

func (api templatePost) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola, este es un servidor HTTP básico en Go!")
}

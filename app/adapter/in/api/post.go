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
	post := templatePost{}
	pattern := http.MethodPost + " " + conf.ApiPrefix +
		"/insert-your-custom-pattern-here"
	log.Println(pattern)
	http.HandleFunc(pattern, post.handler)
	return nil
}

func (api templatePost) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unimplemented")
}

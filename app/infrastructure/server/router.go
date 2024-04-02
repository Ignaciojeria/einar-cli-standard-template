package server

import (
	"archetype/app/configuration"
	"fmt"
	"log"
	"net/http"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

var _ = ioc.Registry(NewRouter, configuration.NewConf)

type Router struct {
	mux                *http.ServeMux
	conf               configuration.Conf
	registeredPatterns []string
}

func NewRouter(conf configuration.Conf) Router {
	server := Router{
		mux:  http.NewServeMux(),
		conf: conf,
	}
	return server
}

func (r *Router) GET(pattern string, handler http.HandlerFunc) {
	pattern = http.MethodGet + " " + r.conf.ApiPrefix + pattern
	r.registeredPatterns = append(r.registeredPatterns, pattern)
	r.mux.Handle(pattern, handler)
}

func (r *Router) POST(pattern string, handler http.HandlerFunc) {
	pattern = http.MethodPost + " " + r.conf.ApiPrefix + pattern
	r.registeredPatterns = append(r.registeredPatterns, pattern)
	r.mux.Handle(pattern, handler)
}

func (r *Router) PATCH(pattern string, handler http.HandlerFunc) {
	pattern = http.MethodPatch + " " + r.conf.ApiPrefix + pattern
	r.registeredPatterns = append(r.registeredPatterns, pattern)
	r.mux.Handle(pattern, handler)
}

func (r *Router) PUT(pattern string, handler http.HandlerFunc) {
	pattern = http.MethodPut + " " + r.conf.ApiPrefix + pattern
	r.registeredPatterns = append(r.registeredPatterns, pattern)
	r.mux.Handle(pattern, handler)
}

func (r *Router) DELETE(pattern string, handler http.HandlerFunc) {
	pattern = http.MethodDelete + " " + r.conf.ApiPrefix + pattern
	r.registeredPatterns = append(r.registeredPatterns, pattern)
	r.mux.Handle(pattern, handler)
}

func (r Router) PrintPatterns() {
	fmt.Println("registered patterns :")
	for _, v := range r.registeredPatterns {
		fmt.Println(v)
	}
}

func (r Router) ServeHTTP() {
	r.PrintPatterns()
	log.Println("Listening on port :" + r.conf.Port)
	log.Fatal(http.ListenAndServe(":"+r.conf.Port, r.mux))
}

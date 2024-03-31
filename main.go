package main

import (
	_ "archetype/app/configuration"
	"log"
	"net/http"
	"os"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

func main() {
	if err := ioc.LoadDependencies(); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("Port")
	log.Println("starting server at port :" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("error starting server : ", err)
	}
}

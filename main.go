package main

import (
	"archetype/app/shared/constants"
	_ "archetype/app/shared/infrastructure/healthcheck"
	_ "archetype/app/shared/infrastructure/observability"
	"archetype/app/shared/infrastructure/serverwrapper"
	_ "embed"
	"log"
	"os"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

//go:embed .version
var version string

func main() {
	os.Setenv(constants.Version, version)
	if err := ioc.LoadDependencies(); err != nil {
		log.Fatal(err)
	}
	serverwrapper.Start()
}

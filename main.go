package main

import (
	"archetype/app/shared/configuration"
	"archetype/app/shared/constants"
	_ "archetype/app/shared/infrastructure/healthcheck"
	"archetype/app/shared/infrastructure/observability"
	"archetype/app/shared/infrastructure/serverwrapper"
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"go.opentelemetry.io/otel"
)

//go:embed .version
var version string

func main() {
	os.Setenv(constants.Version, version)
	if err := ioc.LoadDependencies(); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	tp, err := observability.NewTracerProvider(ctx,
		ioc.Get[configuration.Conf](configuration.NewConf))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			fmt.Println(err.Error())

		}
	}(ctx)
	serverwrapper.Start()
}

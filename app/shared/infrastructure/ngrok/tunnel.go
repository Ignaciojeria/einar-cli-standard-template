package ngrok

import (
	"archetype/app/shared/infrastructure/serverwrapper"
	"context"

	"time"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func init() {
	ioc.Registry(newTunnel, serverwrapper.NewEchoWrapper)
}
func newTunnel(w serverwrapper.EchoWrapper) error {
	ctx, cancel := context.WithCancel(context.Background())
	var success bool
	go func() {
		time.Sleep(10 * time.Second)
		if !success {
			cancel()
		}
	}()
	tunel, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err == nil {
		success = true
	}
	w.Echo.Listener = tunel
	return err
}

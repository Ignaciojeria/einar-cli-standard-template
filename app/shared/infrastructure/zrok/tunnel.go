package ngrok

import (
	"archetype/app/shared/infrastructure/serverwrapper"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/openziti/zrok/environment"
	"github.com/openziti/zrok/sdk/golang/sdk"
)

func init() {
	ioc.Registry(newTunnel, serverwrapper.NewEchoWrapper)
}
func newTunnel(w serverwrapper.EchoWrapper) error {
	root, err := environment.LoadRoot()
	if err != nil {
		return err
	}
	shr, err := sdk.CreateShare(root, &sdk.ShareRequest{
		BackendMode: sdk.ProxyBackendMode,
		ShareMode:   sdk.PublicShareMode,
		Frontends:   []string{"public"},
		Target:      "http-server",
	})

	if err != nil {
		return err
	}

	listenner, err := sdk.NewListener(shr.Token, root)
	if err != nil {
		return err
	}
	w.Echo.Listener = listenner
	return nil
}

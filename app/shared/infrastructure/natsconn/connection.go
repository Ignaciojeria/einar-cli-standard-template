package nats

import (
	"archetype/app/shared/configuration"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/nats-io/nats.go"
)

func init() {
	ioc.Registry(New, configuration.NewNatsConfiguration)
}
func New(conf configuration.NatsConfiguration) (*nats.Conn, error) {
	return nats.Connect(
		conf.NATS_CONNECTION_URL,
		nats.UserCredentials(conf.NATS_CONNECTION_CREDS_FILEPATH))
}

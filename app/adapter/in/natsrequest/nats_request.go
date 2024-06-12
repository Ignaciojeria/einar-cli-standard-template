package natsrequest

import (
	"archetype/app/shared/infrastructure/natsconn"
	"archetype/app/shared/logging"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/nats-io/nats.go"
)

func init() {
	ioc.Registry(newNatsRequest, natsconn.NewConn, logging.NewLogger)
}
func newNatsRequest(nc *nats.Conn, logger logging.Logger) (*nats.Subscription, error) {
	return nc.Subscribe("help.request", func(msg *nats.Msg) {
		nc.Publish(msg.Reply, []byte("Here is the help you requested"))
	})
}

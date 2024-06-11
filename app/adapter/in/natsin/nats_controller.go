package natsin

import (
	"archetype/app/shared/infrastructure/natsconn"
	"archetype/app/shared/logging"
	"context"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/nats-io/nats.go/jetstream"
)

func init() {
	ioc.Registry(NewNatsController, natsconn.NewJetStream, logging.NewLogger)
}
func NewNatsController(js jetstream.JetStream, logger logging.Logger) (jetstream.ConsumeContext, error) {
	ctx := context.Background()
	consumer, err := js.CreateOrUpdateConsumer(ctx, "stream-name", jetstream.ConsumerConfig{
		Name:          "consumer-name",
		Durable:       "consumer-name",
		MaxAckPending: 5,
	})
	if err != nil {
		return nil, err
	}
	return consumer.Consume(func(msg jetstream.Msg) {
		logger.Info("Received message", "data", string(msg.Data()))
		msg.Ack()
	})
}

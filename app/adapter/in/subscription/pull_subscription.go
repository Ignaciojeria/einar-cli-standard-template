package subscription

import (
	"archetype/app/exception"
	pubsubWrapper "archetype/app/infrastructure/pubsub"
	"archetype/app/infrastructure/pubsub/subscription"
	"archetype/app/infrastructure/server"
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/pubsub"
	ioc "github.com/Ignaciojeria/einar-ioc"
)

func init() {
	ioc.Registry(
		NewPullSubscription,
		pubsubWrapper.NewClientWrapper,
		server.NewEchoWrapper)
}
func NewPullSubscription(
	client pubsubWrapper.ClientWrapper,
	httpServer server.EchoWrapper) {
	subscriptionName := "INSERT_YOUR_SUBSCRIPTION_NAME_HERE"
	subRef := client.Subscription(subscriptionName)
	subRef.ReceiveSettings.MaxOutstandingMessages = 5
	settings := subRef.Receive
	go subscription.
		New(subscriptionName, Pull, settings).
		WithPushHandler("/subscription/"+subscriptionName, httpServer).
		Start()
}

func Pull(
	ctx context.Context,
	subscriptionName string,
	m *pubsub.Message) (statusCode int, err error) {

	var dataModel interface{}
	defer func() {
		subscription.HandleMessageAcknowledgement(ctx, &subscription.HandleMessageAcknowledgementDetails{
			MessageID:        m.ID,
			PublishTime:      m.PublishTime.String(),
			SubscriptionName: subscriptionName,
			Error:            err,
			Message:          m,
			ErrorsRequiringNack: []error{
				exception.INTERNAL_SERVER_ERROR,
				exception.EXTERNAL_SERVER_ERROR,
				exception.HTTP_NETWORK_ERROR,
				exception.PUBSUB_BROKER_ERROR,
			},
			CustomLogFields: map[string]interface{}{
				"dataModel": dataModel,
			},
		})
	}()
	if err := json.Unmarshal(m.Data, &dataModel); err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

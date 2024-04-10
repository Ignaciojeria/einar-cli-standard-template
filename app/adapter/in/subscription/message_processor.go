package subscription

import (
	"archetype/app/exception"
	"archetype/app/infrastructure/pubsubwrapper"
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/pubsub"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

type MessageProcessor struct {
	subscriptionReference *pubsub.Subscription
}

func init() {
	ioc.Registry(
		NewMessageProcessor,
		pubsubwrapper.NewSubscriptionManager)
}

func NewMessageProcessor(
	subscriptionManager pubsubwrapper.SubscriptionManager) pubsubwrapper.MessageProcessor {
	subscriptionName := "INSERT_YOUR_SUBSCRIPTION_NAME_HERE"
	subscriptionRef := subscriptionManager.Subscription(subscriptionName)
	subscriptionRef.ReceiveSettings.MaxOutstandingMessages = 5
	messageProcessor := MessageProcessor{
		subscriptionReference: subscriptionRef,
	}
	go subscriptionManager.
		WithMessageProcessor(messageProcessor).
		WithPushHandler("/subscription/" + subscriptionName).
		Start()
	return messageProcessor
}

func (p MessageProcessor) Pull(ctx context.Context, m *pubsub.Message) (statusCode int, err error) {
	var dataModel interface{}
	defer func() {
		pubsubwrapper.HandleMessageAcknowledgement(ctx, &pubsubwrapper.HandleMessageAcknowledgementDetails{
			MessageID:        m.ID,
			PublishTime:      m.PublishTime.String(),
			SubscriptionName: p.subscriptionReference.String(),
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

func (p MessageProcessor) SubscriptionRef() *pubsub.Subscription {
	return p.subscriptionReference
}

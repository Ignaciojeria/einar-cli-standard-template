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

func (p messageProcessorStruct) subscriptionName() string {
	return "INSERT_YOUR_SUBSCRIPTION_NAME_HERE"
}
func (p messageProcessorStruct) Pull(ctx context.Context, m *pubsub.Message) (statusCode int, err error) {
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
		return http.StatusNoContent, err
	}
	return http.StatusOK, nil
}

/* ----- Default Initialization & Configuration Settings ----- */

type messageProcessorStruct struct {
	subscriptionReference *pubsub.Subscription
}

func init() {
	ioc.Registry(
		newMessageProcessor,
		pubsubwrapper.NewSubscriptionManager)
}

func newMessageProcessor(
	subscriptionManager pubsubwrapper.SubscriptionManager) pubsubwrapper.MessageProcessor {
	messageProcessor := messageProcessorStruct{}
	subscriptionRef := subscriptionManager.Subscription(messageProcessor.subscriptionName())
	subscriptionRef.ReceiveSettings.MaxOutstandingMessages = 5
	messageProcessor.subscriptionReference = subscriptionRef
	sm := subscriptionManager.
		WithMessageProcessor(messageProcessor).
		WithPushHandler("/subscription/" + messageProcessor.subscriptionName())
	go sm.Start()
	return messageProcessor
}

func (p messageProcessorStruct) SubscriptionRef() *pubsub.Subscription {
	return p.subscriptionReference
}

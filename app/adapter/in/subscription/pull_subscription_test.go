package subscription

import (
	"archetype/app/configuration"
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
	ioc "github.com/Ignaciojeria/einar-ioc"
)

func TestPull(t *testing.T) {
	defer ioc.NewMockBehaviourForTesting[configuration.Conf](configuration.NewConf,
		configuration.Conf{}).Release()
	ctx := context.Background()
	testData := `{"key": "value"}`
	msg := &pubsub.Message{
		Data: []byte(testData),
	}
	_, err := Pull(ctx, "fake-subscription-dont-change", msg)
	if err != nil {
		t.Errorf("archetype_subscription returned an error: %v", err)
	}
}

func TestArchetypeSubscriptionInvalidInput(t *testing.T) {
	ctx := context.Background()
	invalidTestData := `invalid json`
	msg := &pubsub.Message{
		Data: []byte(invalidTestData),
	}
	_, err := Pull(ctx, "fake-subscription-dont-change", msg)
	if err == nil {
		t.Errorf("archetype_subscription did not return an error for invalid input")
	}
}

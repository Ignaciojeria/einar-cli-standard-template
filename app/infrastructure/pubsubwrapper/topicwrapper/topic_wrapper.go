package topicwrapper

import (
	"archetype/app/infrastructure/pubsubwrapper"
	"sync"

	"cloud.google.com/go/pubsub"
	ioc "github.com/Ignaciojeria/einar-ioc"
)

type TopicWrapper struct {
	clientWrapper pubsubwrapper.ClientWrapper
	topicRefs     sync.Map
}

func init() {
	ioc.Registry(NewTopicWrapper, pubsubwrapper.NewClientWrapper)
}

func NewTopicWrapper(clientWrapper pubsubwrapper.ClientWrapper) *TopicWrapper {
	return &TopicWrapper{
		clientWrapper: clientWrapper,
	}
}

func Get(topicName string) *pubsub.Topic {
	wrapper := ioc.Get[*TopicWrapper](NewTopicWrapper)
	value, ok := wrapper.topicRefs.Load(topicName)
	if ok {
		// If the topic reference was found, return it.
		return value.(*pubsub.Topic)
	}
	// If the topic reference was not found, create a new one.
	newTopicRef := wrapper.clientWrapper.Client().Topic(topicName)
	// Store the new topic reference in the map.
	wrapper.topicRefs.Store(topicName, newTopicRef)
	// Return the new topic reference.
	return newTopicRef
}

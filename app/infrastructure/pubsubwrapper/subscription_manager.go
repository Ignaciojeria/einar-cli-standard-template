package pubsubwrapper

import (
	"archetype/app/adapter/out/slog"
	"archetype/app/constants"
	"archetype/app/infrastructure/serverwrapper"
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

type SubscriptionManager interface {
	Subscription(id string) *pubsub.Subscription
	WithMessageProcessor(mp MessageProcessor) SubscriptionManager
	WithPushHandler(path string) SubscriptionManager
	Start() (SubscriptionManager, error)
}

type SubscriptionWrapper struct {
	clientWrapper    ClientWrapper
	httpServer       serverwrapper.EchoWrapper
	messageProcessor MessageProcessor
}

const subscription_name = "subscription_name"

func init() {
	ioc.Registry(
		NewSubscriptionManager,
		NewClientWrapper,
		serverwrapper.NewEchoWrapper,
	)
}

func NewSubscriptionManager(cw ClientWrapper, s serverwrapper.EchoWrapper) SubscriptionManager {
	return &SubscriptionWrapper{clientWrapper: cw, httpServer: s}
}

func newSubscriptionManagerWithMessageProcessor(
	cw ClientWrapper,
	s serverwrapper.EchoWrapper,
	mp MessageProcessor) SubscriptionManager {
	return &SubscriptionWrapper{clientWrapper: cw, httpServer: s, messageProcessor: mp}
}

func (sw *SubscriptionWrapper) Subscription(id string) *pubsub.Subscription {
	return sw.clientWrapper.Subscription(id)
}

func (sw SubscriptionWrapper) WithMessageProcessor(mp MessageProcessor) SubscriptionManager {
	return newSubscriptionManagerWithMessageProcessor(sw.clientWrapper, sw.httpServer, mp)
}

func (s *SubscriptionWrapper) Start() (SubscriptionManager, error) {
	ctx := context.Background()
	if err := s.messageProcessor.SubscriptionRef().Receive(ctx, s.receive); err != nil {
		slog.Logger().Error(
			"subscription_signal_broken",
			subscription_name, s.messageProcessor.SubscriptionRef().String(),
			constants.Error, err.Error(),
		)
		time.Sleep(10 * time.Second)
		go s.Start()
		return s, err
	}
	return s, nil
}

func (s *SubscriptionWrapper) receive(ctx context.Context, m *pubsub.Message) {
	s.messageProcessor.Pull(ctx, m)
}

func (s *SubscriptionWrapper) WithPushHandler(path string) SubscriptionManager {
	s.httpServer.POST(path, s.pushHandler)
	return s
}

func (s *SubscriptionWrapper) pushHandler(c echo.Context) error {
	googleChannel := c.Request().Header.Get("X-Goog-Channel-ID")

	var msg pubsub.Message

	if googleChannel != "" {
		if err := c.Bind(&msg); err != nil {
			return c.String(http.StatusNoContent, "error binding Pub/Sub message")
		}
		statusCode, err := s.messageProcessor.Pull(c.Request().Context(), &msg)
		if statusCode >= 500 && statusCode <= 599 {
			return c.String(statusCode, "")
		}
		if err != nil {
			return c.String(http.StatusNoContent, "")
		}
		return c.String(http.StatusOK, "")
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "error reading request body")
	}

	msg.Attributes = make(map[string]string)
	for key, values := range c.Request().Header {
		if len(values) > 0 {
			msg.Attributes[strings.ToLower(key)] = strings.Join(values, ",")
		}
	}

	msg.Data = body
	if statusCode, err := s.messageProcessor.Pull(c.Request().Context(), &msg); err != nil {
		return c.String(statusCode, err.Error())
	}
	return c.String(http.StatusOK, "")

}

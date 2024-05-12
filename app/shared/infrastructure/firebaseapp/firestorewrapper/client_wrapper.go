package firestorewrapper

import (
	"archetype/app/shared/infrastructure/firebaseapp"

	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	ioc "github.com/Ignaciojeria/einar-ioc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientWrapper struct {
	client *firestore.Client
}

func init() {
	ioc.Registry(NewClientWrapper, firebaseapp.NewFirebaseAPP)
}

func NewClientWrapper(app *firebase.App) (*ClientWrapper, error) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		return &ClientWrapper{}, err
	}
	_, err = client.Collection("health").Doc("ping").Get(ctx)
	if status.Code(err) != codes.NotFound {
		return nil, err
	}
	return &ClientWrapper{
		client: client,
	}, nil
}

func (wrapper *ClientWrapper) Client() *firestore.Client {
	return wrapper.client
}

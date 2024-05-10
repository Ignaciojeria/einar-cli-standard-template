package firestorewrapper

import (
	"archetype/app/shared/infrastructure/firebaseapp"

	"context"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	ioc "github.com/Ignaciojeria/einar-ioc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientWrapper struct {
	collectionRefs sync.Map
	client         *firestore.Client
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

func (wrapper *ClientWrapper) Collection(collectionName string) *firestore.CollectionRef {
	value, ok := wrapper.collectionRefs.Load(collectionName)
	if ok {
		// If the collection reference was found, return it.
		return value.(*firestore.CollectionRef)
	}
	// If the collection reference was not found, create a new one.
	newCollectionRef := wrapper.client.Collection(collectionName)
	// Store the new collection reference in the map.
	wrapper.collectionRefs.Store(collectionName, newCollectionRef)
	// Return the new collection reference.
	return newCollectionRef
}

func (wrapper *ClientWrapper) Client() *firestore.Client {
	return wrapper.client
}

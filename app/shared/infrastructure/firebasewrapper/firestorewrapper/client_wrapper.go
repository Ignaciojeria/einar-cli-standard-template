package firestorewrapper

import (
	"archetype/app/shared/constants"
	"archetype/app/shared/infrastructure/firebasewrapper"
	"archetype/app/shared/logger"

	"context"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	ioc "github.com/Ignaciojeria/einar-ioc"
)

type ClientWrapper struct {
	collectionRefs sync.Map
	client         *firestore.Client
}

func init() {
	ioc.Registry(NewClientWrapper, firebasewrapper.NewFirebaseAPP)
}

func NewClientWrapper(app *firebase.App) (*ClientWrapper, error) {
	client, err := app.Firestore(context.Background())
	if err != nil {
		logger.Logger().Error("error getting firestore client", constants.Error, err.Error())
		return &ClientWrapper{}, err
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

package firestorewrapper

import (
	"archetype/app/adapter/out/slog"
	"archetype/app/constants"
	"archetype/app/infrastructure/firebasewrapper"
	"context"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	ioc "github.com/Ignaciojeria/einar-ioc"
	"go.opentelemetry.io/otel"
)

type ClientWrapper struct {
	collectionRefs sync.Map
	client         *firestore.Client
}

var Tracer = otel.Tracer("firestorewrapper")

func init() {
	ioc.Registry(NewClientWrapper, firebasewrapper.NewFirebaseAPP)
}

func NewClientWrapper(app *firebase.App) (ClientWrapper, error) {
	client, err := app.Firestore(context.Background())
	if err != nil {
		slog.Logger().Error("error getting firestore client", constants.Error, err.Error())
		return ClientWrapper{}, err
	}
	return ClientWrapper{
		client: client,
	}, nil
}

func Collection(collectionName string) *firestore.CollectionRef {
	wrapper := ioc.Get[*ClientWrapper](NewClientWrapper)
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

func GetFirestoreClient(collectionName string) *firestore.Client {
	wrapper := ioc.Get[*ClientWrapper](NewClientWrapper)
	return wrapper.client
}

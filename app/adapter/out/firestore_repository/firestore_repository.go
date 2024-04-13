package firestoredb

import "archetype/app/infrastructure/firebasewrapper/firestorewrapper"

func RunFirestoreOperation() {
	var _ = firestorewrapper.
		GetFirestoreCollection("insert your collection path here")
}

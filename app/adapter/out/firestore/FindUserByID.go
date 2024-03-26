package firestore

import (
	"context"
	"errors"
	"log"
	"os"
	einar "te-eme-backend/app/shared/archetype/firestore"
	"te-eme-backend/app/shared/archetype/model"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/api/iterator"
)

var FindUserByID = func(ctx context.Context, userMap map[string]string) (model.User, error) {
	_, span := einar.Tracer.Start(ctx,
		"Login",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	userQuery := einar.Collection(os.Getenv("Firestore_UsersCollection")).Where("ID", "==", userMap["ID"])

	iter := userQuery.Documents(ctx)

	var userWasFounded bool
	var user model.User

	for {

		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error al obtener documento: %v", err)
			continue
		}

		doc.DataTo(&user)
		userWasFounded = true

	}

	if userWasFounded {

		return user, nil

	} else {
		return model.User{}, errors.New(" [Firestore] : User not founded")
	}

}

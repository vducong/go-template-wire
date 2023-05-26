package databases

import (
	"context"
	"fmt"
	"go-template-wire/configs"
	"go-template-wire/pkg/failure"
	"go-template-wire/pkg/logger"

	"cloud.google.com/go/firestore"
)

type FirestoreDB = *firestore.Client

func NewFirestoreDB(cfg *configs.Config, log *logger.Logger) (FirestoreDB, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, cfg.GCP.ProjectID)
	if err != nil {
		return nil, failure.ErrWithTrace(fmt.Errorf("Failed to init Firestore: %w", err))
	}

	log.Info("Firestore connection established")
	return client, nil
}

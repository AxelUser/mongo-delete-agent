package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AdminHelper struct {
	admin *mongo.Database
}

func CreateAdminHelper(ctx context.Context, uri string) (*AdminHelper, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	admin := client.Database("admin")
	return &AdminHelper{
		admin: admin,
	}, nil
}

func (h AdminHelper) IsCommandRunning(ctx context.Context, jobId int64) (bool, error) {
	cur, err := h.admin.Aggregate(ctx, bson.A{
		bson.D{{Key: "$currentOp", Value: bson.D{{Key: "localOps", Value: true}}}},
		bson.D{{Key: "$match", Value: bson.D{{Key: "command.comment", Value: fmt.Sprintf("job:%d", jobId)}}}},
		bson.D{{Key: "$limit", Value: 1}},
	})

	if err != nil {
		return false, fmt.Errorf("failed to check if command is running: %w", err)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		return true, nil
	}

	return false, nil
}

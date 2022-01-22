package seed

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(uri string, db string, col string) (err error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	defer func() {
		if err != nil {
			err = fmt.Errorf("failed to seed mongo: %w", err)
		}
	}()

	if err != nil {
		return err
	}

	err = create(client, db, col)
	if err != nil {
		return err
	}

	return nil
}

func create(c *mongo.Client, db string, col string) error {
	dbs, err := c.ListDatabaseNames(context.Background(), bson.D{})
	if err != nil {
		return fmt.Errorf("failed to get list of databases: %w", err)
	}

	if contains(dbs, db) {
		err = c.Database(db).Drop(context.Background())
		if err != nil {
			return fmt.Errorf("failed to drop database '%s': %w", db, err)
		}
	}

	err = c.Database(db).CreateCollection(context.Background(), col)
	if err != nil {
		return fmt.Errorf("failed to create collection '%s.%s': %w", db, col, err)
	}

	err = c.Database("admin").RunCommand(context.Background(), bson.D{
		{Key: "enableSharding", Value: db},
	}).Err()

	if err != nil {
		return fmt.Errorf("failed to shard db '%s': %w", db, err)
	}

	err = c.Database("admin").RunCommand(context.Background(), bson.D{
		{Key: "shardCollection", Value: fmt.Sprintf("%s.%s", db, col)},
		{Key: "key", Value: bson.D{{Key: "userId", Value: 1}}},
	}).Err()

	if err != nil {
		return fmt.Errorf("failed to shard collection '%s.%s': %w", db, col, err)
	}

	return nil
}

func contains(l []string, t string) bool {
	for _, v := range l {
		if v == t {
			return true
		}
	}

	return false
}

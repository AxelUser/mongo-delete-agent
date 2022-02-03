package seed

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AxelUser/mongo-delete-agent/src/constants"
	"github.com/AxelUser/mongo-delete-agent/src/entities"
	"github.com/AxelUser/mongo-delete-agent/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(ctx context.Context, conf Config) (err error) {
	log.Printf("Starting seeding '%s.%s' at '%s'", conf.Db, conf.Col, conf.Uri)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Uri))
	defer func() {
		if err != nil {
			err = fmt.Errorf("failed to seed mongo: %w", err)
		}

		if client != nil {
			dErr := client.Disconnect(ctx)
			if dErr != nil {
				// ignore what happend before
				err = fmt.Errorf("failed to disconnect from mongo: %w", dErr)
			}
		}
	}()

	if err != nil {
		return err
	}

	crCol, err := createCol(ctx, client, conf.Db, conf.Col)
	if err != nil {
		return err
	}

	err = seedRndData(ctx, crCol, conf.Accounts, conf.Users, conf.Events)
	if err != nil {
		return err
	}

	err = shardCol(ctx, client, crCol)
	if err != nil {
		return err
	}

	return nil
}

func createCol(ctx context.Context, c *mongo.Client, db string, col string) (*mongo.Collection, error) {
	dbs, err := c.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to get list of databases: %w", err)
	}

	if utils.Contains(dbs, db) {
		err = c.Database(db).Drop(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to drop database '%s': %w", db, err)
		}
		log.Printf("Dropped exising collection '%s.%s'", db, col)
	}

	err = c.Database(db).CreateCollection(ctx, col)
	if err != nil {
		return nil, fmt.Errorf("failed to create collection '%s.%s': %w", db, col, err)
	}

	log.Printf("Created collection '%s.%s'", db, col)

	return c.Database(db).Collection(col), nil
}

func seedRndData(ctx context.Context, col *mongo.Collection, clientCount int, usersCount int, eventsCount int) error {
	log.Printf("Starting seeding '%s.%s' with %d random entities", col.Database().Name(), col.Name(), clientCount*usersCount*eventsCount)
	var insCount int64
	buffer := make([]interface{}, eventsCount)
	for clId := 0; clId < clientCount; clId++ {
		for usr := 0; usr < usersCount; usr++ {
			for e := 0; e < len(buffer); e++ {
				buffer[e] = entities.Event{
					ClientId: int64(clId),
					UserId:   int64(usr),
					TypeId:   constants.AllEventTypes[e%len(constants.AllEventTypes)],
					Props: map[string]string{
						"value": fmt.Sprintf("Value for user %d", clId),
					},
					Time: time.Now().UTC(),
				}
			}
			res, err := col.InsertMany(ctx, buffer)
			if err != nil {
				return fmt.Errorf("failed to insert documents: %w", err)
			}
			insCount += int64(len(res.InsertedIDs))
		}
	}

	log.Printf("Inserted %d events", insCount)

	return nil
}

func shardCol(ctx context.Context, c *mongo.Client, col *mongo.Collection) error {
	index := bson.D{
		{Key: "clientId", Value: 1},
		{Key: "userId", Value: 1},
	}

	_, err := col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: index,
	})
	if err != nil {
		return fmt.Errorf("failed to create index for sharded key")
	}

	log.Printf("Created index {clientId: 1, userId: 1} for '%s.%s'", col.Database().Name(), col.Name())

	err = c.Database("admin").RunCommand(ctx, bson.D{
		{Key: "enableSharding", Value: col.Database().Name()},
	}).Err()

	if err != nil {
		return fmt.Errorf("failed to shard db '%s': %w", col.Database().Name(), err)
	}

	log.Printf("Enabled sharding for database '%s'", col.Database().Name())

	err = c.Database("admin").RunCommand(ctx, bson.D{
		{Key: "shardCollection", Value: fmt.Sprintf("%s.%s", col.Database().Name(), col.Name())},
		{Key: "key", Value: index}}).Err()

	if err != nil {
		return fmt.Errorf("failed to shard collection '%s.%s': %w", col.Database().Name(), col.Name(), err)
	}

	log.Printf("Sharded collection '%s.%s'", col.Database().Name(), col.Name())
	return nil
}

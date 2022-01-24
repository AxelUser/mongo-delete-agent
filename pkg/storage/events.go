package storage

import (
	"context"
	"fmt"

	"github.com/AxelUser/mongo-delete-agent/pkg/config"
	"github.com/AxelUser/mongo-delete-agent/pkg/entities"
	"github.com/AxelUser/mongo-delete-agent/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventsRepository struct {
	col *mongo.Collection
}

func CreateEventsRepository(ctx context.Context, conn config.MongoConnection) (*EventsRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn.Uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	col := client.Database(conn.Db).Collection(conn.Col)
	return &EventsRepository{
		col: col,
	}, nil
}

type AccountStats struct {
	AccountId int64 `bson:"_id"`
	Events    int64 `bson:"events"`
}

func (r *EventsRepository) CountByAccount(ctx context.Context) ([]AccountStats, error) {
	cur, err := r.col.Aggregate(ctx, bson.A{
		bson.D{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$clientId"},
			{Key: "events", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}},
		bson.D{{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to count events by accounts: %w", err)
	}

	defer cur.Close(ctx)

	var stats []AccountStats
	if err := cur.All(ctx, &stats); err != nil {
		return nil, fmt.Errorf("failed to decode account statistics: %w", err)
	}

	return stats, nil
}

func (r *EventsRepository) Get(ctx context.Context, c models.ClientId, u models.UserId, t models.EventTypeId, skip int64, take int64) ([]entities.Event, error) {
	cur, err := r.col.Find(ctx, bson.D{
		{Key: "clientId", Value: c},
		{Key: "userId", Value: u},
		{Key: "typeId", Value: t},
	}, &options.FindOptions{
		Skip:  &skip,
		Limit: &take,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get events for 'ClientId:%d, UserId:%d, TypeId:%s': %w", c, u, t, err)
	}

	var res []entities.Event
	if err := cur.All(ctx, &res); err != nil {
		return nil, fmt.Errorf("failed to get events for 'ClientId:%d, UserId:%d, TypeId:%s': %w", c, u, t, err)
	}

	return res, nil
}

func (r *EventsRepository) Exists(ctx context.Context, q models.DataQuery) (bool, error) {
	res := r.col.FindOne(ctx, q.GetFilter())
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, fmt.Errorf("failed to find events for query '%s': %w", q, err)
	}

	return true, nil
}

func (r *EventsRepository) InsertMany(ctx context.Context, ls []interface{}) (int, error) {
	res, err := r.col.InsertMany(ctx, ls)
	if err != nil {
		return 0, fmt.Errorf("failed to insert %d events: %w", len(ls), err)
	}

	return len(res.InsertedIDs), nil
}

func (r *EventsRepository) Delete(ctx context.Context, q models.DataQuery) (int64, error) {
	res, err := r.col.DeleteMany(ctx, q.GetFilter())
	if err != nil {
		return 0, fmt.Errorf("failed to remove events for query '%s': %w", q, err)
	}

	return res.DeletedCount, nil
}

func (r *EventsRepository) Close(ctx context.Context) error {
	if r.col != nil {
		err := r.col.Database().Client().Disconnect(ctx)
		if err != nil {
			// ignore what happend before
			return fmt.Errorf("failed to disconnect from mongo: %w", err)
		}
	}
	return nil
}

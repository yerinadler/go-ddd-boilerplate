package infrastructure

import (
	"context"

	"github.com/yerinadler/go-ddd/internal/domain/order"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COLLECTION_NAME = "orders"
)

type MongoOrderRepository struct {
	db *mongo.Collection
}

func NewMongoOrderRepository(db *mongo.Database) order.OrderRepository {
	coll := db.Collection(COLLECTION_NAME)
	return &MongoOrderRepository{
		db: coll,
	}
}

func (repo *MongoOrderRepository) Save(ctx context.Context, entity *order.Order) error {
	count, err := repo.db.CountDocuments(ctx, bson.M{"id": entity.Id})

	if err != nil {
		return err
	}

	if count > 0 {
		_, err = repo.db.ReplaceOne(ctx, bson.M{"id": entity.Id}, entity)
		if err != nil {
			return err
		}
		return nil
	}

	_, err = repo.db.InsertOne(ctx, entity)

	if err != nil {
		return err
	}

	return nil
}

func (repo *MongoOrderRepository) GetById(ctx context.Context, id string) (*order.Order, error) {
	var order *order.Order
	err := repo.db.FindOne(ctx, bson.M{"id": id}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

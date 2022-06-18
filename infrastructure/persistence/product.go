package persistence

import (
	"context"
	"fmt"

	"github.com/danisbagus/simple-go-ddd/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "products"

type ProductRepository struct {
	coll *mongo.Collection
}

func NewProductRepo(db *mongo.Database) ProductRepository {
	return ProductRepository{db.Collection(collectionName)}
}

func (r ProductRepository) Insert(ctx context.Context, product *entity.Product) error {
	res, err := r.coll.InsertOne(ctx, product)
	if err != nil {
		return fmt.Errorf("failed insert product: %v", err)
	}

	if res.InsertedID == "" {
		return fmt.Errorf("failed insert product: no data was inserted")
	}

	return nil
}

func (r ProductRepository) FindAll() ([]entity.Product, error) {
	products := make([]entity.Product, 0)
	filter := bson.M{}

	cursor, err := r.coll.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed get list product: %v", err)
	}

	err = cursor.All(context.Background(), &products)
	if err != nil {
		return nil, fmt.Errorf("failed read all cursor: %v", err)
	}

	return products, nil
}

func (r ProductRepository) FindOneByID(ID string) (*entity.Product, error) {
	oid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, fmt.Errorf("failed convert object id: %v", err)
	}

	product := new(entity.Product)
	filter := bson.M{"_id": oid}
	res := r.coll.FindOne(context.Background(), filter)
	if err := res.Decode(product); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("failed get product: data not found")

		}
		return nil, fmt.Errorf("failed get product: %v", err)
	}

	return product, nil
}

func (r ProductRepository) Update(ctx context.Context, ID string, product *entity.Product) error {
	oid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return fmt.Errorf("failed convert object id: %v", err)
	}

	filter := bson.M{"_id": oid}

	result, err := r.coll.UpdateOne(ctx, filter, bson.M{"$set": product})
	if err != nil {
		return fmt.Errorf("failed update product: %v", err)
	}

	if result.ModifiedCount < 1 {
		return fmt.Errorf("failed update product: no data was updated")
	}

	return nil
}

func (r ProductRepository) Delete(ctx context.Context, ID string) error {
	oid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return fmt.Errorf("failed convert object id: %v", err)
	}

	res, err := r.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return fmt.Errorf("failed delete product: %v", err)
	}

	if res.DeletedCount < 1 {
		return fmt.Errorf("failed delete product: no data was deleted")
	}

	return nil
}

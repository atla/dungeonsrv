package repository

import (
	"context"
	"errors"
	"time"

	"github.com/atla/dungeonsrv/pkg/db"
	"github.com/atla/dungeonsrv/pkg/entities"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type itemsRepository struct {
	db         *db.Client
	collection string
}

//NewMongoItemsRepository creates a new mongodb ItemsRepository
func NewMongoItemsRepository(db *db.Client) ItemsRepository {
	return &itemsRepository{
		db:         db,
		collection: "items",
	}
}

func (ir *itemsRepository) FindByID(id string) (*entities.Item, error) {

	result := ir.db.FindByID(ir.collection, id)

	if result != nil {

		item := &entities.Item{}
		if err := result.Decode(item); err != nil {
			log.WithField("Error", err).Error("Error decoding search object")
			return nil, errors.New("Search not found")
		}
		return item, nil
	}
	return nil, errors.New("Item not found")
}

// FindAll returns all series
func (ir *itemsRepository) FindAll() ([]*entities.Item, error) {

	var results []*entities.Item
	cursor, err := ir.db.FindAll(ir.collection)

	if err != nil {
		log.WithField("collection", ir.collection).WithField("cursor", cursor).Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var elem entities.Item
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	return results, nil
}

// Store stores a new item
func (ir *itemsRepository) Store(item *entities.Item) (*entities.Item, error) {

	// update creation time
	item.Created = time.Now()

	if result, error := ir.db.InsertOne(ir.collection, item); error != nil {
		log.WithField("Error", error).Error("error during insertion")
		return nil, error
	} else {

		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			item.ID = oid
		}
		return item, nil
	}
}

// Store stores a new item
func (ir *itemsRepository) Update(item *entities.Item) error {

	if result, error := ir.db.UpdateOneByID(ir.collection, item.ID, item); error != nil {
		log.WithField("Error", error).Error("error during insertion")
		return error
	} else {
		log.WithField("Items", result).Info("updated item")
	}
	return nil
}

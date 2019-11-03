package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ItemType type
type ItemType int

const (
	itemTypeCurrency ItemType = iota + 1
	itemTypeConsumable
	itemTypeArmor
	itemTypeWeapon
	itemTypeCollectible
)

//Item data
type Item struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
	ItemType    ItemType           `bson:"itemType,omitempty" json:"itemType"`
	Created     time.Time          `bson:"created,omitempty" json:"created,omitempty"`
}

//Items type
type Items []*Item

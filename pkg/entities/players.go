package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Player data
type Player struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
	LastSeen    time.Time          `bson:"lastSeen,omitempty" json:"lastSeen,omitempty"`
	Created     time.Time          `bson:"itemType,omitempty" json:"created,omitempty"`
}

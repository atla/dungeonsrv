package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayerStatus int8

const (
	playerStatusOnline  PlayerStatus = 1
	playerStatusOffline PlayerStatus = 2
)

//Player data
type Player struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
	LastSeen    time.Time          `bson:"lastSeen,omitempty" json:"lastSeen,omitempty"`
	Created     time.Time          `bson:"created,omitempty" json:"created,omitempty"`
	Status      PlayerStatus       `bson:"playerStatus,omitempty" json:"playerStatus,omitempty"`
}

package entities

type RoomActionType int8

const (
	RoomActionTypeDirection RoomActionType = iota + 1
)

type RoomAction struct {
}

//Room data
type Room struct {
	ID          string `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string `bson:"name,omitempty" json:"name"`
	Description string `bson:"description,omitempty" json:"description"`
}

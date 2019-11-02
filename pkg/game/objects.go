package game

type ObjectType int

const (
	ObjectTypePlayer   ObjectType = 0
	ObjectTypeNPC      ObjectType = 1
	ObjectTypeItem     ObjectType = 2
	ObjectTypeResource ObjectType = 3
)

//
type Object struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

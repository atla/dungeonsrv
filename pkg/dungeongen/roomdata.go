package dungeongen

import (
	"errors"
	"log"
)

//RoomData ...
type RoomData struct {
	X      int
	Y      int
	Width  int
	Height int

	IsConnected bool
	Visited     bool
	Section     uint8
	doors       []RoomDoor
}

//NewRoomData creates a new room data instance
func NewRoomData(x int, y int, width int, height int) *RoomData {
	return &RoomData{
		X:           x,
		Y:           y,
		Width:       width,
		Height:      height,
		IsConnected: false,
		Visited:     false,
		doors:       nil,
	}
}

//RoomDoor... room door
type RoomDoor struct {
	Direction int
	Position  Vec2D
}

//NewRoomDoor ... creates a new room door
func NewRoomDoor(direction int, pos Vec2D) RoomDoor {

	if direction < 0 || direction > 3 {
		log.Fatal("room direction not between 0 and 4")
	}
	return RoomDoor{
		Direction: direction,
		Position:  pos,
	}
}

//GetWallForPosition ...
func (r *RoomData) GetWallForPosition(x, y int) (int, error) {

	// west wall
	if x == r.X && y >= r.Y && y <= (r.Y+r.Height) {
		return DirectionWest, nil
	}
	// north wall
	if y == r.Y && x >= r.X && x <= (r.X+r.Width) {
		return DirectionNorth, nil
	}
	// east wall
	if x == (r.X+r.Width) && y >= r.Y && y <= (r.Y+r.Height) {
		return DirectionEast, nil
	}
	// south wall
	if y == (r.Y+r.Height) && x >= r.X && x <= (r.X+r.Width) {
		return DirectionSouth, nil
	}
	return -1, errors.New("Position not on wall")

}

// HasDoor ...returns if room has at least one door
func (r *RoomData) HasDoor(direction int) bool {
	for _, door := range r.doors {
		if door.Direction == direction {
			return true
		}
	}
	return false
}

// AddDoor ...
func (r *RoomData) AddDoor(direction int, pos Vec2D) {

	r.doors = append(r.doors, NewRoomDoor(direction, pos))
}

// Doors ...
func (r *RoomData) Doors() []RoomDoor {
	return r.doors
}

//IsCorner returns if coord is a room corner (dont add doors there)
func (r *RoomData) IsCorner(x, y int) bool {

	if (r.Y == y && r.X == x) || (r.Y == y && r.X+r.Width == x) || (r.Y+r.Height == y && r.X == x) || (r.Y+r.Height == y && r.X+r.Width == x) {
		return true
	}
	return false
}

// Collides returns true if two rooms overlap
func (r *RoomData) Collides(r2 RoomData) bool {
	if r.X < r2.X+r2.Width &&
		r.X+r.Width > r2.X &&
		r.Y < r2.Y+r2.Height &&
		r.Y+r.Height > r2.Y {
		return true
	}
	return false
}

// IsInside returns true if a point is within the bounds of the room
func (r *RoomData) IsInside(x, y int) bool {
	return x >= r.X && x <= r.X+r.Width && y >= r.Y && y <= r.Y+r.Height
}

// Extrude extrudes a room by factor returning a bigger or smaller room
func (r *RoomData) Extrude(factor int) *RoomData {
	return &RoomData{
		X:           r.X - factor,
		Y:           r.Y - factor,
		Width:       r.Width + (factor * 2),
		Height:      r.Height + (factor * 2),
		IsConnected: r.IsConnected,
		doors:       r.Doors(),
	}
}

package dungeongen

//RoomData ...
type RoomData struct {
	X      int
	Y      int
	Width  int
	Height int

	IsConnected bool
}

//NewRoomData creates a new room data instance
func NewRoomData(x int, y int, width int, height int) *RoomData {
	return &RoomData{
		X:           x,
		Y:           y,
		Width:       width,
		Height:      height,
		IsConnected: false,
	}
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

// Extrude extrudes a room by factor returning a bigger or smaller room
func (r *RoomData) Extrude(factor int) *RoomData {
	return &RoomData{
		X:      r.X - factor,
		Y:      r.Y - factor,
		Width:  r.Width + (factor * 2),
		Height: r.Height + (factor * 2),
	}
}

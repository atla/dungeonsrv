package dungeongen

import (
	"math/rand"
	"time"
)

// RandomRoomStrategy...
type RandomRoomStrategy struct {
	Density           RoomDensity
	MaxRooms          int
	UseRandomSeed     bool
	Seed              int
	MaxRoomWidth      int
	MaxRoomHeight     int
	MinRoomWidth      int
	MinRoomHeight     int
	SpaceBetweenRooms int
}

type RoomDensity int8

const (
	RoomDensityLow    RoomDensity = 1
	RoomDensityMedium             = 2
	RoomDensityBig                = 3
	RoomDensityMax                = 4
)

// NewRandomRoomStrategy returns a default RandomRoomStrategy
func NewRandomRoomStrategy() *RandomRoomStrategy {
	return &RandomRoomStrategy{
		Density:           RoomDensityMedium,
		MaxRooms:          -1,
		UseRandomSeed:     true,
		MinRoomWidth:      10,
		MinRoomHeight:     10,
		MaxRoomWidth:      50,
		MaxRoomHeight:     50,
		SpaceBetweenRooms: 2,
	}
}

const (
	DirectionWest  = 0
	DirectionNorth = 1
	DirectionEast  = 2
	DirectionSouth = 3
)

func getMaxRoomsForDensity(data *DungeonData, density RoomDensity) int {
	baseFor100x100 := float32(int(density) * 50.0)
	factor := float32(((data.Width + data.Height) / 2) / 100.0)
	return int(baseFor100x100 * factor)
}

// Create ...
func (strategy *RandomRoomStrategy) Create(data *DungeonData) {

	// update max rooms
	if strategy.MaxRooms < 0 {
		strategy.MaxRooms = getMaxRoomsForDensity(data, strategy.Density)
	}

	// 1st Step: Create rooms
	for i := 0; i < strategy.MaxRooms; i++ {
		newRoom := strategy.createRandomRoom(data)

		if !strategy.roomCollidesWithExisting(data, newRoom) {
			addRoomToDungeon(data, newRoom)
		}
	}

	// 2nd Step: Create hallways between rooms
	for _, room := range data.Rooms {

		if room.IsConnected {
			continue
		}

		// Select a random wall
		start, direction := strategy.selectRandomWall(room)
		collision := false
		current := start

		// create a new door at the starting point
		data.Set(current.X, current.Y, DoorTileType)

		//TODO: add exit/action to room (not roomdata)

		// walk direction until collision

		for !collision {
			next := current.Add(direction)

			//collided with dungeon bounds
			if data.IsOutside(next.X, next.Y) {
				collision = true
				break
			}

			tileType := data.Get(next.X, next.Y)

			if tileType != EmptyTileType {
				collision = true
			} else {
				data.Set(next.X, next.Y, FloorTileType)
			}
			current = next
		}

		switch data.Get(current.X, current.Y) {
		case FloorTileType:
			room.IsConnected = true
			break
		case WallTileType:
			data.Set(current.X, current.Y, DoorTileType)
		case DoorTileType:
			// TODO: find according room
			room.IsConnected = true
			break
		}
	}
}

func (strategy *RandomRoomStrategy) selectRandomWall(room *RoomData) (Vec2D, Vec2D) {

	wall := rand.Int() % 4
	var start Vec2D
	var direction Vec2D

	switch wall {
	case DirectionWest:
		start = NewVec2D(room.X, room.Y+1+(rand.Int()%(room.Height-2)))
		direction = NewVec2D(-1, 0)
		break
	case DirectionNorth:
		start = NewVec2D(room.X+1+(rand.Int()%(room.Width-2)), room.Y)
		direction = NewVec2D(0, -1)
		break
	case DirectionEast:
		start = NewVec2D(room.X+room.Width-1, room.Y+(rand.Int()%(room.Height-2)))
		direction = NewVec2D(1, 0)
		break
	case DirectionSouth:
		start = NewVec2D(room.X+(rand.Int()%(room.Width-2)), room.Y+room.Height-1)
		direction = NewVec2D(0, 1)
		break
	}
	return start, direction
}

func (strategy *RandomRoomStrategy) roomCollidesWithExisting(data *DungeonData, room *RoomData) bool {

	extruded := room.Extrude(1)

	for _, r := range data.Rooms {
		// extrude rooms by 1 so we get some spacing between rooms
		if r.Collides(*extruded) {
			return true
		}
	}
	return false
}

func addRoomToDungeon(data *DungeonData, room *RoomData) {

	data.Rooms = append(data.Rooms, room)

	for x := room.X; x < room.X+room.Width; x++ {
		for y := room.Y; y < room.Y+room.Height; y++ {

			// is wall
			if x == room.X || y == room.Y || x == (room.X+room.Width-1) || y == (room.Y+room.Height-1) {
				data.Set(x, y, WallTileType)
			} else {
				data.Set(x, y, FloorTileType)
			}
		}
	}
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func (strategy *RandomRoomStrategy) createRandomRoom(data *DungeonData) *RoomData {

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	w := max(strategy.MinRoomWidth, r.Int()%strategy.MaxRoomWidth)
	h := max(strategy.MinRoomHeight, r.Int()%strategy.MaxRoomHeight)
	x := max(0, (r.Int()%data.Width - w))
	y := max(0, (r.Int()%data.Height - h))

	roomData := &RoomData{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}

	return roomData
}

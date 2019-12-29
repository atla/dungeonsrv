package dungeongen

import (
	"errors"
	"time"

	"github.com/atla/dungeonsrv/pkg/util"
)

// Explorer is used to explore a generated dungeon and create walkable path to be able to generate an interactive story
type Explorer struct {
}

//NewExplorer convenience function to create an explorer instance
func NewExplorer() *Explorer {
	return &Explorer{}
}

func selectRandomUnvisitedRoom(rooms []*RoomData) (*RoomData, error) {

	for _, room := range rooms {
		if !room.Visited {
			return room, nil
		}
	}
	return nil, errors.New("all rooms visited")
}

func followPath(pos Vec2D, startingRoom *RoomData, data *DungeonData) {

	if data.GetPath(pos.X, pos.Y) == EmptyTileType {

		switch data.Get(pos.X, pos.Y) {
		case FloorTileType:
			// mark and follow
			data.SetPath(pos.X, pos.Y, PathTileType)

			// spawn GoRoutines?
			followPath(NewVec2D(pos.X+1, pos.Y), startingRoom, data)
			followPath(NewVec2D(pos.X-1, pos.Y), startingRoom, data)
			followPath(NewVec2D(pos.X, pos.Y+1), startingRoom, data)
			followPath(NewVec2D(pos.X, pos.Y-1), startingRoom, data)
			break

		case DoorTileType:
			data.SetPath(pos.X, pos.Y, PathTileType)
			// visit new room
			if room, err := data.FindRoomForCoord(pos.X, pos.Y); err == nil {
				visitRoom(data, room)
			}
			break
		}
	}
}

// function to recursively traverse rooms and pathways between rooms
func visitRoom(data *DungeonData, room *RoomData) {

	if !room.Visited {
		room.Visited = true
		data.SetRoomPath(room.X+1, room.Y+1, room.Width-1, room.Height-1, PathTileType)
		//TODO: parallelizee traversal?
		for direction := range room.Doors() {
			doorPos := room.GetDoor(direction)
			data.SetPath(doorPos.X, doorPos.Y, PathTileType)

			switch direction {
			case DirectionWest:
				followPath(NewVec2D(doorPos.X-1, doorPos.Y), room, data)
				break
			case DirectionNorth:
				followPath(NewVec2D(doorPos.X, doorPos.Y-1), room, data)
				break
			case DirectionEast:
				followPath(NewVec2D(doorPos.X+1, doorPos.Y), room, data)
				break
			case DirectionSouth:
				followPath(NewVec2D(doorPos.X, doorPos.Y+1), room, data)
				break
			}

		}
	}
}

// Explore starts the dungeon exploring
func (e *Explorer) Explore(data *DungeonData) {

	defer util.TimeTrack(time.Now(), "Explorer")

	room, err := selectRandomUnvisitedRoom(data.Rooms)

	if err == nil {
		visitRoom(data, room)
	}
	/*
		var room *RoomData
		var err error

		room, err = selectRandomUnvisitedRoom(data.Rooms)

		for err == nil {
			visitRoom(data, room)
			room, err = selectRandomUnvisitedRoom(data.Rooms)
		}*/
}

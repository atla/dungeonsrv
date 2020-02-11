package dungeongen

import (
	"errors"
	"fmt"
	"math/rand"
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

func followPath(section uint8, pos Vec2D, startingRoom *RoomData, data *DungeonData) {

	if data.GetPath(pos.X, pos.Y) == 0 {

		switch data.Get(pos.X, pos.Y) {
		case FloorTileType:
			// mark and follow
			data.SetPath(pos.X, pos.Y, section)

			// spawn GoRoutines?
			followPath(section, NewVec2D(pos.X+1, pos.Y), startingRoom, data)
			followPath(section, NewVec2D(pos.X-1, pos.Y), startingRoom, data)
			followPath(section, NewVec2D(pos.X, pos.Y+1), startingRoom, data)
			followPath(section, NewVec2D(pos.X, pos.Y-1), startingRoom, data)
			break

		case DoorTileType:
			data.SetPath(pos.X, pos.Y, section)
			// visit new room
			if room, err := data.FindRoomForCoord(pos.X, pos.Y); err == nil {
				visitRoom(section, data, room)
			}

			// spawn GoRoutines?
			/*followPath(section, NewVec2D(pos.X+1, pos.Y), startingRoom, data)
			followPath(section, NewVec2D(pos.X-1, pos.Y), startingRoom, data)
			followPath(section, NewVec2D(pos.X, pos.Y+1), startingRoom, data)
			followPath(section, NewVec2D(pos.X, pos.Y-1), startingRoom, data)
			*/
			break
		default:

			break

		}
	}
}

// function to recursively traverse rooms and pathways between rooms
func visitRoom(section uint8, data *DungeonData, room *RoomData) {

	if !room.Visited {
		room.Visited = true
		room.Section = section

		data.SetRoomPath(room.X+1, room.Y+1, room.Width-1, room.Height-1, section)

		//TODO: parallelizee traversal?
		for _, door := range room.Doors() {
			doorPos := door.Position
			data.SetPath(doorPos.X, doorPos.Y, section)

			switch door.Direction {
			case DirectionWest:
				followPath(section, NewVec2D(doorPos.X-1, doorPos.Y), room, data)
				break
			case DirectionNorth:
				followPath(section, NewVec2D(doorPos.X, doorPos.Y-1), room, data)
				break
			case DirectionEast:
				followPath(section, NewVec2D(doorPos.X+1, doorPos.Y), room, data)
				break
			case DirectionSouth:
				followPath(section, NewVec2D(doorPos.X, doorPos.Y+1), room, data)
				break
			default:
				fmt.Println("VISITROOM door without direction, this should not happen")
			}

		}
	}
}

// Explore starts the dungeon exploring
func (e *Explorer) Explore(data *DungeonData) {

	defer util.TimeTrack(time.Now(), "Explorer")

	var room *RoomData
	var err error
	var section uint8
	section = 1

	room, err = selectRandomUnvisitedRoom(data.Rooms)

	for err == nil {
		visitRoom(section, data, room)
		room, err = selectRandomUnvisitedRoom(data.Rooms)
		section++
		//fmt.Println(err)
	}

	// reset visited status of all rooms
	for _, room := range data.Rooms {
		room.Visited = false
	}

	if section > 1 {

		var sectionsConnected uint8
		var currentSection uint8
		sectionsConnected = 1
		currentSection = 1

		for sectionsConnected < section-1 {
			var result *intersectingRoomResult
			var originRoom *RoomData

			if rooms, err := getRoomsOfSection(currentSection, data); err == nil {

				foundRoom := false

				for _, room := range rooms {

					fmt.Printf("Room of section %d\n", room.Section)

					for _, direction := range randomDirectionArray() {
						//find connecting room of different
						var err error
						result, err = tryFindIntersectingRoom(direction, outwardDirectionForWall(direction), room, currentSection, data)

						if err == nil {
							originRoom = room
							foundRoom = true
							break
						}
					}

					if foundRoom {
						break
					}
				}

				//try building a pathway between the two rooms
				if foundRoom && buildPathwayBetweenSectionRooms(originRoom, result, data) {

					removeSection := result.Room.Section

					// mark all rooms of the newly connected section
					data.ForEachTile(func(x, y int, tile TileType, data *DungeonData) {
						// transform all tiles of section to currentsection
						if tile != EmptyTileType && data.GetPath(x, y) == removeSection {
							data.SetPath(x, y, currentSection)
						}
					})

					if roomsOfSection, err := getRoomsOfSection(removeSection, data); err == nil {
						for _, sroom := range roomsOfSection {
							sroom.Section = currentSection
						}
					}

					sectionsConnected += 2
					currentSection++
				}
			}
		}

		fmt.Println("Finished connecting all sections")
	}
}

func buildPathwayBetweenSectionRooms(origin *RoomData, result *intersectingRoomResult, data *DungeonData) bool {

	stop := result.CollisionPos

	for pos := result.StartingPos; !(pos.X == stop.X && pos.Y == stop.Y); pos = NewVec2D(pos.X+result.Direction.X, pos.Y+result.Direction.Y) {
		data.Set(pos.X, pos.Y, FloorTileType)
	}
	//add two doors
	data.Set(result.StartingPos.X, result.StartingPos.Y, DoorTileType)
	data.Set(result.CollisionPos.X, result.CollisionPos.Y, DoorTileType)

	wallDirection, _ := origin.GetWallForPosition(result.StartingPos.X, result.StartingPos.Y)
	origin.AddDoor(wallDirection, result.StartingPos)

	wallDirectionCollision, _ := result.Room.GetWallForPosition(result.StartingPos.X, result.StartingPos.Y)
	result.Room.AddDoor(wallDirectionCollision, result.StartingPos)

	return true
}

func outwardDirectionForWall(direction int) Vec2D {
	if direction == DirectionWest {
		return NewVec2D(-1, 0)
	}
	if direction == DirectionNorth {
		return NewVec2D(0, -1)
	}
	if direction == DirectionEast {
		return NewVec2D(1, 0)
	}
	return NewVec2D(0, 1)
}

func randomDirectionArray() (res [4]int) {
	start := rand.Int() % 4
	for i := 0; i < 4; i++ {
		res[i] = (start + i) % 4
	}
	return
}

type intersectingRoomResult struct {
	Room                  *RoomData
	StartingPos           Vec2D
	CollisionPos          Vec2D
	StartingWallDirection int
	Direction             Vec2D
}

func tryFindIntersectingRoom(wallDirection int, direction Vec2D, room *RoomData, section uint8, data *DungeonData) (*intersectingRoomResult, error) {

	var startingPos Vec2D

	switch wallDirection {
	case DirectionWest:
		startingPos = NewVec2D(room.X, room.Y+1+(rand.Int()%(room.Height-2)))
		break
	case DirectionNorth:
		startingPos = NewVec2D(room.X+1+(rand.Int()%(room.Width-2)), room.Y)
		break
	case DirectionEast:
		startingPos = NewVec2D(room.X+room.Width, room.Y+1+(rand.Int()%(room.Height-2)))
		break
	case DirectionSouth:
		startingPos = NewVec2D(room.X+1+(rand.Int()%(room.Width-2)), room.Y+room.Height)
		break
	}

	currentPos := startingPos
	collided := false
	//foundRoom := false
	for !collided {

		currentPos = NewVec2D(currentPos.X+direction.X, currentPos.Y+direction.Y)

		if data.IsOutside(currentPos.X, currentPos.Y) {
			return nil, errors.New("Outside of bounds")
		}

		// check if we found a room
		if collidedRoom, err := data.FindRoomForCoord(startingPos.X, startingPos.Y); err == nil {

			// check if collision was with wall
			if _, err := collidedRoom.GetWallForPosition(startingPos.X, startingPos.Y); err != nil {
				return nil, errors.New("Could not find room")
			}

			if collidedRoom.Section > 0 && collidedRoom.Section != section {
				return &intersectingRoomResult{
					Room:                  collidedRoom,
					StartingPos:           startingPos,
					CollisionPos:          currentPos,
					StartingWallDirection: wallDirection,
					Direction:             direction,
				}, nil
			}

			return nil, errors.New("Did collide with same section room")
		}
	}

	return nil, errors.New("Did not find an intersecting room with different section")
}

func filter(ss []*RoomData, filterFunc func(*RoomData) bool) (ret []*RoomData) {
	for _, s := range ss {
		if filterFunc(s) {
			ret = append(ret, s)
		}
	}
	return
}

func getRoomsOfSection(section uint8, data *DungeonData) ([]*RoomData, error) {
	// TODO: select rooms that are most west/north/east/south

	testFunc := func(r *RoomData) bool {
		return r.Section == section
	}

	result := filter(data.Rooms, testFunc)

	if len(result) == 0 {
		return nil, fmt.Errorf("No Room Found with section %d", section)
	}
	return result, nil
}

// deprecated?
func selectRandomRoomOfSection(section uint8, data *DungeonData) (*RoomData, error) {

	// TODO: select rooms that are most west/north/east/south

	testFunc := func(r *RoomData) bool {
		return r.Section == section
	}

	result := filter(data.Rooms, testFunc)

	if len(result) == 0 {
		return nil, errors.New("No Room Found with section")
	} else {

		num := rand.Int() % len(result)
		return result[num], nil
	}
}

package dungeongen

import "errors"

//DungeonCreationStrategy ...
type DungeonCreationStrategy interface {
	Create(data *DungeonData, mask AreaMask)
}

//Builder ..
type Builder interface {
	Build() *DungeonData

	WithSmallSize() Builder
	WithMask(mask AreaMask) Builder
	WithSize(width int, height int) Builder
	WithCreationStrategy(strategy DungeonCreationStrategy) Builder
}

//DungeonData ...
type DungeonData struct {
	Width  int
	Height int

	MapData  MapData2D
	PathData PathData2D
	Rooms    []*RoomData
}

//TileType ...
type TileType int16

const (
	EmptyTileType TileType = iota + 1
	FloorTileType
	WallTileType
	DoorTileType
	//PathTileType
)

//MapData2D ...
type MapData2D []TileType

//PathData2D ...
type PathData2D []uint8

// IsOutside checks if coords are outside of the dungeon space
func (data *DungeonData) IsOutside(x, y int) bool {
	return x < 0 || y < 0 || x > data.Width-1 || y > data.Height-1
}

func (data *DungeonData) ForEachTile(fun func(x, y int, tile TileType, data *DungeonData)) {

	for x := 0; x < data.Width; x++ {
		for y := 0; y < data.Height; y++ {
			fun(x, y, data.Get(x, y), data)
		}
	}
}

//Set ...
func (data *DungeonData) Set(x int, y int, tile TileType) {

	// ignore Set va.ues out of bounds
	if x < 0 || x >= data.Width || y < 0 || y >= data.Height {
		return
	}
	data.MapData[x+y*data.Width] = tile
}

//SetPath ...
func (data *DungeonData) SetPath(x int, y int, tile uint8) {

	// ignore Set values out of bounds
	if x < 0 || x >= data.Width || y < 0 || y >= data.Height {
		return
	}
	data.PathData[x+y*data.Width] = tile
}

//SetRoomPath ...
func (data *DungeonData) SetRoomPath(x, y, width, height int, tile uint8) {

	// ignore Set va.ues out of bounds
	if x < 0 || (x+width) > data.Width || y < 0 || (y+height) > data.Height {
		return
	}
	for xp := x; xp < x+width; xp++ {
		for yp := y; yp < y+height; yp++ {
			data.PathData[xp+yp*data.Width] = tile
		}
	}
}

//FindRoomForCoord ...
func (data *DungeonData) FindRoomForCoord(x, y int) (*RoomData, error) {
	for _, room := range data.Rooms {
		if room.IsInside(x, y) {
			return room, nil
		}
	}
	return nil, errors.New("Could not find Room at coord")
}

//Get ...
func (data *DungeonData) Get(x int, y int) TileType {

	// ignore Set values out of bounds
	if x < 0 || x >= data.Width || y < 0 || y >= data.Height {
		return -1
	}
	return data.MapData[x+y*data.Width]
}

//GetPath ...
func (data *DungeonData) GetPath(x int, y int) uint8 {

	// ignore Set values out of bounds
	if x < 0 || x >= data.Width || y < 0 || y >= data.Height {
		return 0
	}
	return data.PathData[x+y*data.Width]
}

//Init ...
func (data *DungeonData) Init() {
	data.MapData = make([]TileType, data.Width*data.Height)
	data.PathData = make([]uint8, data.Width*data.Height)

	for x := 0; x < data.Width; x++ {
		for y := 0; y < data.Height; y++ {
			data.Set(x, y, EmptyTileType)
			data.SetPath(x, y, 0)
		}
	}
}

type defaultBuilder struct {
	Data     *DungeonData
	Strategy DungeonCreationStrategy
	Mask     AreaMask
}

//DefaultBuilder ...
func DefaultBuilder() Builder {
	builder := &defaultBuilder{}
	builder.Mask = &EmptyMask{}
	builder.Data = &DungeonData{
		Width:  100,
		Height: 100,
	}
	builder.Strategy = &RandomRoomStrategy{
		MaxRooms:      10,
		MinRoomWidth:  10,
		MinRoomHeight: 10,
	}
	builder.Data.Width = 100
	builder.Data.Height = 100
	return builder
}

func (builder *defaultBuilder) WithMask(mask AreaMask) Builder {
	builder.Mask = mask
	return builder
}

func (builder *defaultBuilder) WithCreationStrategy(strategy DungeonCreationStrategy) Builder {
	builder.Strategy = strategy
	return builder
}

func (builder *defaultBuilder) WithSmallSize() Builder {
	builder.Data.Width = 120
	builder.Data.Height = 120
	return builder
}

func (builder *defaultBuilder) WithSize(width int, height int) Builder {
	builder.Data.Width = width
	builder.Data.Height = height
	return builder
}

//Build
func (builder *defaultBuilder) Build() *DungeonData {

	builder.Data.Init()

	// Invoke strategy
	builder.Strategy.Create(builder.Data, builder.Mask)

	/*
		&CircleMask{
			Radius:  (builder.Data.Width / 2) - builder.Data.Width/10,
			CenterX: builder.Data.Width / 2,
			CenterY: builder.Data.Height / 2,
		}
	*/

	explorer := NewExplorer()
	explorer.Explore(builder.Data)

	return builder.Data
}

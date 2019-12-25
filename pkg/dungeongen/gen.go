package dungeongen

//DungeonCreationStrategy ...
type DungeonCreationStrategy interface {
	Create(data *DungeonData)
}

//DungeonData ...
type DungeonData struct {
	Width  int
	Height int

	MapData MapData2D
	Rooms   []*RoomData
}

//TileType ...
type TileType int16

//MapData2D ...
type MapData2D []TileType

func (data *DungeonData) IsOutside(x, y int) bool {
	return x < 0 || y < 0 || x > data.Width-1 || y > data.Height-1
}

//Set ...
func (data *DungeonData) Set(x int, y int, tile TileType) {

	// ignore Set va.ues out of bounds
	if x < 0 || x > data.Width {
		return
	}
	// ignore Set va.ues out of bounds
	if y < 0 || y > data.Height {
		return
	}

	data.MapData[x+y*data.Width] = tile
}

//Get ...
func (data *DungeonData) Get(x int, y int) TileType {

	// ignore Set va.ues out of bounds
	if x < 0 || x > data.Width {
		return -1
	}
	// ignore Set va.ues out of bounds
	if y < 0 || y > data.Height {
		return -1
	}

	return data.MapData[x+y*data.Width]
}

const (
	EmptyTileType TileType = iota + 1
	FloorTileType
	WallTileType
	DoorTileType
)

//Init ...
func (data *DungeonData) Init() {
	data.MapData = make([]TileType, data.Width*data.Height)

	for x := 0; x < data.Width; x++ {
		for y := 0; y < data.Height; y++ {
			data.Set(x, y, EmptyTileType)
		}
	}
}

//Builder ..
type Builder interface {
	Build() *DungeonData
	WithSmallSize() Builder
	WithSize(width int, height int) Builder
	WithCreationStrategy(strategy DungeonCreationStrategy) Builder
}

type defaultBuilder struct {
	Data     *DungeonData
	Strategy DungeonCreationStrategy
}

//DefaultBuilder ...
func DefaultBuilder() Builder {
	builder := &defaultBuilder{}
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
	builder.Strategy.Create(builder.Data)

	return builder.Data
}

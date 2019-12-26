package dungeongen

const (
	DirectionWest  = 0
	DirectionNorth = 1
	DirectionEast  = 2
	DirectionSouth = 3
)

type DirectionCallback struct {
	DirectionCallbackWest  func()
	DirectionCallbackNorth func()
	DirectionCallbackEast  func()
	DirectionCallbackSouth func()
}

func (dc *DirectionCallback) On(direction int) {
	switch direction {
	case 0:
		dc.DirectionCallbackWest()
		return
	case 1:
		dc.DirectionCallbackNorth()
		return
	case 2:
		dc.DirectionCallbackEast()
		return
	case 3:
		dc.DirectionCallbackSouth()
		return
	}
}

func NewDirectionCallback(west, north, east, south func()) *DirectionCallback {
	return &DirectionCallback{
		DirectionCallbackWest:  west,
		DirectionCallbackNorth: north,
		DirectionCallbackEast:  east,
		DirectionCallbackSouth: south,
	}
}

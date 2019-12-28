package dungeongen

type Vec2D struct {
	X int
	Y int
}

func NewVec2D(x, y int) Vec2D {
	return Vec2D{
		X: x,
		Y: y,
	}
}

func (v Vec2D) Add(v2 Vec2D) Vec2D {
	return Vec2D{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vec2D) Invert() Vec2D {
	return Vec2D{
		X: -v.X,
		Y: -v.Y,
	}
}

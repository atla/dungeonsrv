package main

import (
	"fmt"

	"github.com/atla/dungeonsrv/pkg/dungeongen"
)

func main() {

	rs := dungeongen.NewRandomRoomStrategy()
	rs.MaxRoomWidth = 20
	rs.MaxRoomHeight = 20
	rs.Density = dungeongen.RoomDensityMax

	gen := dungeongen.DefaultBuilder().WithCreationStrategy(rs).WithSize(150, 150)

	dungeon := gen.Build()

	fmt.Printf("Created dungeon. size %dx%d", dungeon.Width, dungeon.Height)

	exporter := dungeongen.PNGExporter{}
	exporter.Export(*dungeon, dungeongen.FormatPNG)

}

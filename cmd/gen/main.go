package main

import (
	"fmt"

	"github.com/atla/dungeonsrv/pkg/dungeongen"
)

func main() {

	gen := dungeongen.DefaultBuilder().WithCreationStrategy(dungeongen.NewRandomRoomStrategy()).WithSize(200, 200)
	dungeon := gen.Build()

	fmt.Printf("Created dungeon. size %dx%d", dungeon.Width, dungeon.Height)

	exporter := dungeongen.PNGExporter{}
	exporter.Export(*dungeon, dungeongen.FormatPNG)

}

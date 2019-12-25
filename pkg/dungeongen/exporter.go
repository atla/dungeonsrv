package dungeongen

import (
	"image"
	"image/color"
	"image/png"

	"github.com/nfnt/resize"

	"log"
	"os"
)

//ExporterFormat is the format used for exporter the DungeonData
type ExporterFormat int8

const (
	FormatPNG ExporterFormat = iota + 1
	FormatJSON
)

// Exporter exports data according to format
type Exporter interface {
	Export(data DungeonData, format ExporterFormat) interface{}
}

type PNGExporter struct {
}

var floorTileColor color.NRGBA = color.NRGBA{
	R: 255,
	G: 128,
	B: 128,
	A: 255,
}
var doorTileColor color.NRGBA = color.NRGBA{
	R: 80,
	G: 200,
	B: 80,
	A: 255,
}

var wallTileColor color.NRGBA = color.NRGBA{
	R: 255,
	G: 200,
	B: 200,
	A: 255,
}

var emptyTileColor color.NRGBA = color.NRGBA{
	R: 0,
	G: 0,
	B: 0,
	A: 255,
}

// Export PNG
func (exp *PNGExporter) Export(data DungeonData, format ExporterFormat) interface{} {

	width := data.Width
	height := data.Height

	img := image.NewNRGBA(image.Rectangle{Max: image.Point{X: width, Y: height}})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch data.Get(x, y) {
			case FloorTileType:
				img.SetNRGBA(x, y, floorTileColor)
				break
			case WallTileType:
				img.SetNRGBA(x, y, wallTileColor)
				break
			case DoorTileType:
				img.SetNRGBA(x, y, doorTileColor)
				break
			default:
				//img.SetNRGBA(x, y, emptyTileColor)
				break
			}
		}
	}

	img2 := resize.Resize(uint(data.Width*2), 0, img.SubImage(img.Rect), resize.NearestNeighbor)

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img2); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	return img
}

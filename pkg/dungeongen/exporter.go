package dungeongen

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"time"

	"log"
	"os"

	"github.com/nfnt/resize"
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

//TODO: Extract to utils package?
//ParseHexColor ...
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func GetRandomColorInRgb() (c color.RGBA) {

	rand.Seed(time.Now().UnixNano())
	c.R = uint8(rand.Intn(255))
	c.G = uint8(rand.Intn(255))
	c.B = uint8(rand.Intn(255))
	c.A = 255

	return c
}

// Export PNG
func (exp *PNGExporter) ExportAsImage(data DungeonData, format ExporterFormat) *image.Image {

	theme := make(map[TileType]color.RGBA)
	theme[FloorTileType], _ = ParseHexColor("#3B4252")
	theme[DoorTileType], _ = ParseHexColor("#81A1C1")
	theme[WallTileType], _ = ParseHexColor("#4C566A")
	theme[EmptyTileType], _ = ParseHexColor("#2E3440")
	//theme[PathTileType], _ = ParseHexColor("#33aa33")

	sectionColors := make(map[uint8]color.RGBA)

	width := data.Width
	height := data.Height

	img := image.NewRGBA(image.Rectangle{Max: image.Point{X: width, Y: height}})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			tile := data.Get(x, y)

			if color, ok := theme[tile]; ok {
				img.SetRGBA(x, y, color)
			}

			pathTile := data.GetPath(x, y)
			if pathTile > 0 {
				if color, ok := sectionColors[pathTile]; ok {
					img.SetRGBA(x, y, color)
				} else {
					c := GetRandomColorInRgb()
					sectionColors[pathTile] = c
					img.SetRGBA(x, y, c)
				}
			}
		}
	}

	img2 := resize.Resize(uint(data.Width*5), 0, img.SubImage(img.Rect), resize.NearestNeighbor)
	return &img2
}

// Export PNG
func (exp *PNGExporter) ExportAsFile(data DungeonData, format ExporterFormat, fileName string) interface{} {

	img2 := exp.ExportAsImage(data, format)

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, *img2); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	return img2
}

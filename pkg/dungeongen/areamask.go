package dungeongen

import (
	"image/png"
	"log"
	"os"
)

//AreaMask ...
type AreaMask interface {
	IsInside(x, y int) bool
}

//EmptyMask ...
type EmptyMask struct {
}

//IsInside tet for EmptyMask
func (em *EmptyMask) IsInside(x, y int) bool {
	return true
}

//CircleMask ...
type CircleMask struct {
	Radius  int
	CenterX int
	CenterY int
}

//IsInside tet for CircleMask
func (cm *CircleMask) IsInside(x, y int) bool {
	return (x-cm.CenterX)*(x-cm.CenterX)+(y-cm.CenterY)*(y-cm.CenterY) < cm.Radius*cm.Radius
}

type BitmapMask struct {
	Width  int
	Height int
	Mask   []bool
}

//check if mask is white at that point
func (bm *BitmapMask) IsInside(x, y int) bool {

	if x < 0 || x > bm.Width-1 || y < 0 || y > bm.Height-1 {
		return false
	}

	return bm.Mask[x+y*bm.Width]
}

//LoadFromFile ...
func LoadFromFile(file string) *BitmapMask {
	// Read image from file that already exists
	existingImageFile, err := os.Open(file)
	if err != nil {
		log.Fatal("Error reading file ", file)
	}
	defer existingImageFile.Close()

	// Alternatively, since we know it is a png
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		log.Fatal("Error decoding file ", err)
	}

	w := loadedImage.Bounds().Size().X
	h := loadedImage.Bounds().Size().Y

	mask := make([]bool, w*h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			r, _, _, _ := loadedImage.At(x, y).RGBA()

			if r == 0 {
				mask[x+y*w] = false
			} else {
				mask[x+y*w] = true
			}
		}
	}

	return &BitmapMask{
		Mask:   mask,
		Width:  loadedImage.Bounds().Size().X,
		Height: loadedImage.Bounds().Size().Y,
	}
}

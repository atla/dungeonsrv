package main

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/atla/dungeonsrv/pkg/dungeongen"
)

func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func randomDensity() dungeongen.RoomDensity {
	switch rand.Int() % 4 {
	case 0:
		return dungeongen.RoomDensityLow
	case 1:
		return dungeongen.RoomDensityMedium
	case 2:
		return dungeongen.RoomDensityHigh
	}
	return dungeongen.RoomDensityMax
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		rs := dungeongen.NewRandomRoomStrategy()
		rs.MinRoomWidth = 5 + rand.Int()%10
		rs.MaxRoomWidth = rs.MinRoomWidth + rand.Int()%20

		rs.MinRoomHeight = 5 + rand.Int()%10
		rs.MaxRoomHeight = rs.MinRoomHeight + rand.Int()%20

		// smooth out width and height
		rs.MinRoomHeight = (rs.MinRoomWidth + rs.MinRoomHeight) / 2
		rs.MaxRoomHeight = (rs.MaxRoomHeight + rs.MaxRoomHeight) / 2

		rs.Density = randomDensity()
		rs.SpaceBetweenRooms = 1 + rand.Int()%3
		rs.ChanceOfAdjacentRooms = 10 + rand.Int()%50
		rs.ChanceForDivergence = 5 + rand.Int()%20

		width := 200  //(2 + rand.Int()%8) * 50
		height := 100 //2 + rand.Int()%8) * 50

		gen := dungeongen.DefaultBuilder().WithCreationStrategy(rs).WithSize(width, height)

		dungeon := gen.Build()

		exporter := dungeongen.PNGExporter{}
		img := exporter.ExportAsImage(*dungeon, dungeongen.FormatPNG)

		writeImage(w, img)
	})

	log.Fatal(http.ListenAndServe(":8083", nil))

}

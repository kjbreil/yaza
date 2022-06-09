package tilegen

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

var (
	emptyImage = ebiten.NewImage(3, 3)

	// emptySubImage is an internal sub image of emptyImage.
	// Use emptySubImage at DrawTriangles instead of emptyImage in order to avoid bleeding edges.
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

	tileSize = 16
)

func Init() {
	emptyImage.Fill(color.White)
}

// IndexCord returns the coordinates associated with an index in a Dimension
func IndexCord(index int) (x int, y int) {
	sx := (index % tileSize) * tileSize
	sy := (index / tileSize) * tileSize
	return sx, sy
}

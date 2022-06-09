package tilegen

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kjbreil/yaza/assets"
	"image"
)

func Dirt(as []*ebiten.Image, screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	ix, iy := IndexCord(17)

	img := as[assets.SP.I("ground")].SubImage(image.Rect(ix, iy, ix+tileSize, iy+tileSize)).(*ebiten.Image)
	screen.DrawImage(img, op)
}

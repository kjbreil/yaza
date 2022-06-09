package tilegen

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func Person(screen *ebiten.Image, x, y int) {

	var path vector.Path
	xf, yf := float32(x), float32(y)

	// TODO: Add curves
	path.MoveTo(xf, yf)
	path.Arc(xf, yf, 5, 0, 360, vector.Clockwise)

	// path.QuadTo(xf+100, yf, xf+50, yf)

	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}
	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0xdb / float32(0xff)
		vs[i].ColorG = 0x56 / float32(0xff)
		vs[i].ColorB = 0x20 / float32(0xff)
	}
	screen.DrawTriangles(vs, is, emptySubImage, op)
}

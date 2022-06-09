package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type gameMap struct {
	// tiles is a map of z(height),x,y
	tiles         map[int]map[int]map[int]*tile
	offset        coordinates
	mapRange      mapRange
	chunkSize     int
	createdChunks map[int]map[int]struct{}
}

type mapRange struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

// createChunk of a map at the given x,y

func initGameMap() *gameMap {
	// TODO: move to parameters
	// widthF, heightF := 640/ebiten.DeviceScaleFactor(), 480/ebiten.DeviceScaleFactor()
	//
	// width, height := int(widthF), int(heightF)

	var gm gameMap

	gm.offset = coordinates{
		x: 0,
		y: 0,
	}

	gm.chunkSize = 160

	gm.tiles = make(map[int]map[int]map[int]*tile)
	gm.createdChunks = make(map[int]map[int]struct{})

	// err := gm.createChunk(0, 0)
	// if err != nil {
	// 	log.Panicln(err)
	// }

	return &gm
}

func (gm *gameMap) expandMap(x1, y1, x2, y2 int) {
	// z := 0

	for x := x1; x < x2; x += gm.chunkSize {
		for y := y1; y < y2; y += gm.chunkSize {
			if _, ok := gm.createdChunks[x][y]; !ok {
				_ = gm.createChunk(x-x%gm.chunkSize, y-y%gm.chunkSize)

			}
		}
	}
}

func (g *Game) drawMap(screen *ebiten.Image) {
	width, height := screen.Size()

	g.GamMap.expandMap(g.GamMap.offset.x-width, g.GamMap.offset.y-height, (width*2)+g.GamMap.offset.x, (height*2)+g.GamMap.offset.y)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("range: %d,%d,%d,%d\noffset: %d,%d\nw+o: %d,%d", g.GamMap.mapRange.x1, g.GamMap.mapRange.y1, g.GamMap.mapRange.x2, g.GamMap.mapRange.y2, g.GamMap.offset.x, g.GamMap.offset.y, width+g.GamMap.offset.x, height+g.GamMap.offset.y))

	// ebitenutil.DebugPrint(screen, fmt.Sprintf("%d,%d", width, height))
	for x := -16; x < width; x++ {
		for y := -16; y < height; y++ {
			for z := range g.GamMap.tiles {
				t := g.GamMap.tileAtCoord(x+g.GamMap.offset.x, y+g.GamMap.offset.y, z)
				// if y+g.GamMap.offset.y > y2 {
				// 	_ = g.GamMap.createChunk(x, y+g.GamMap.offset.y)
				// }
				if t != nil {
					if t.vector != nil {
						t.vector(screen, t.coordinates.x-g.GamMap.offset.x, t.coordinates.y-g.GamMap.offset.y)
					}
					if t.image != nil {
						t.image(g.Assets, screen, t.coordinates.x-g.GamMap.offset.x, t.coordinates.y-g.GamMap.offset.y)
					}
				}
			}
		}
	}

}

func (gm *gameMap) tileAtCoord(x, y, z int) *tile {
	t, ok := gm.tiles[z][x][y]
	if !ok {
		return nil
	}
	return t
}

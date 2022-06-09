package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
)

type gameMap struct {
	// tiles is a map of x,y,z(height)
	tiles         map[int]map[int]map[int]*tile
	offset        coordinates
	mapRange      mapRange
	chunkSize     int
	createdChunks map[int]map[int]struct{}
}

type tile struct {
	text        string
	coordinates *coordinates
}

type mapRange struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func newTile(text string, x, y int) (*tile, *coordinates) {
	c := &coordinates{
		x: x,
		y: y,
	}

	return &tile{
		text:        text,
		coordinates: c,
	}, c
}

// createChunk of a map at the given x,y
// TODO: clear out the chunk first
func (gm *gameMap) createChunk(px, py int) error {
	if px%gm.chunkSize != 0 || py%gm.chunkSize != 0 {
		return fmt.Errorf("chunk x,y incorrect for chunksize")
	}
	z := 0
	chunkOffsetX, chunkOffsetY := px+gm.chunkSize, py+gm.chunkSize
	for x := px; x < chunkOffsetX; x++ {
		for y := py; y < chunkOffsetY; y++ {
			if x%100 == 0 && y%20 == 0 {
				t, c := newTile("B", x, y)
				if _, ok := gm.tiles[c.x]; !ok {
					gm.tiles[c.x] = make(map[int]map[int]*tile)
				}
				if _, ok := gm.tiles[c.x][c.y]; !ok {
					gm.tiles[c.x][c.y] = make(map[int]*tile)
				}
				// if _, ok := gm.tiles[c.x][c.y][z]; !ok {
				gm.tiles[c.x][c.y][z] = t
				// }
			}
		}
	}

	if px < gm.mapRange.x1 {
		gm.mapRange.x1 = px
	}
	if py < gm.mapRange.y1 {
		gm.mapRange.y1 = py
	}
	if chunkOffsetX > gm.mapRange.x2 {
		gm.mapRange.x2 = chunkOffsetX
	}
	if chunkOffsetY > gm.mapRange.y2 {
		gm.mapRange.y2 = chunkOffsetY
	}

	if _, ok := gm.createdChunks[px]; !ok {
		gm.createdChunks[px] = make(map[int]struct{})
	}

	gm.createdChunks[px][py] = struct{}{}

	return nil
}

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
	z := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			t := g.GamMap.tileAtCoord(x+g.GamMap.offset.x, y+g.GamMap.offset.y, z)
			// if y+g.GamMap.offset.y > y2 {
			// 	_ = g.GamMap.createChunk(x, y+g.GamMap.offset.y)
			// }
			if t != nil {
				text.Draw(screen, t.text, g.Font, t.coordinates.x-g.GamMap.offset.x, t.coordinates.y-g.GamMap.offset.y, colornames.Blanchedalmond)
			}
		}
	}

}

func (gm *gameMap) tileAtCoord(x, y, z int) *tile {
	t, ok := gm.tiles[x][y][z]
	if !ok {
		return nil
	}
	return t
}

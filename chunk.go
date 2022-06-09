package main

import (
	"fmt"
	"github.com/kjbreil/yaza/tilegen"
)

// TODO: clear out the chunk first
func (gm *gameMap) createChunk(px, py int) error {
	if px%gm.chunkSize != 0 || py%gm.chunkSize != 0 {
		return fmt.Errorf("chunk x,y incorrect for chunksize")
	}
	chunkOffsetX, chunkOffsetY := px+gm.chunkSize, py+gm.chunkSize
	for x := px; x < chunkOffsetX; x++ {
		for y := py; y < chunkOffsetY; y++ {
			gm.makeChunkGround(x, y, 0)
			gm.makeChunkPlants(x, y, 1)
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
func (gm *gameMap) addTile(t *tile, x, y, z int) {
	if _, ok := gm.tiles[z]; !ok {
		gm.tiles[z] = make(map[int]map[int]*tile)
	}
	if _, ok := gm.tiles[z][x]; !ok {
		gm.tiles[z][x] = make(map[int]*tile)
	}
	gm.tiles[z][x][y] = t
}

func (gm *gameMap) makeChunkGround(x, y, z int) {
	if x%16 == 0 && y%16 == 0 {
		t, c := newImageTile(tilegen.Dirt, x, y, z)
		gm.addTile(t, c.x, c.y, z)
	}

}

func (gm *gameMap) makeChunkPlants(x, y, z int) {
	if x%100 == 0 && y%20 == 0 {
		t, c := newVectorTile(tilegen.Bush, x, y, z)
		gm.addTile(t, c.x, c.y, z)
	}
}

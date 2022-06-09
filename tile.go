package main

import "github.com/hajimehoshi/ebiten/v2"

type tile struct {
	vector      func(screen *ebiten.Image, x, y int)
	image       func(assets []*ebiten.Image, screen *ebiten.Image, x, y int)
	coordinates *coordinates
	depth       int
}

func newVectorTile(v func(screen *ebiten.Image, x, y int), x, y, z int) (*tile, *coordinates) {
	c := &coordinates{
		x: x,
		y: y,
	}

	return &tile{
		vector:      v,
		coordinates: c,
		depth:       z,
	}, c
}

func newImageTile(i func(assets []*ebiten.Image, screen *ebiten.Image, x, y int), x, y, z int) (*tile, *coordinates) {
	c := &coordinates{
		x: x,
		y: y,
	}

	return &tile{
		image:       i,
		coordinates: c,
		depth:       z,
	}, c
}

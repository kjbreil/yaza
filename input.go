package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) personMove() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		g.GamMap.offset.y -= 1
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		g.GamMap.offset.x += 1
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		g.GamMap.offset.y += 1
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		g.GamMap.offset.x -= 1
	}

}

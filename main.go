package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/kjbreil/yaza/resources/statik"
	"github.com/kjbreil/yaza/tilegen"
	"log"
)

func (g *Game) Update() error {

	g.personMove()

	return nil
}

func (g *Game) Init() {
	var err error

	width, height := 640, 480

	tilegen.Init()

	fx, fy := float64(width)/2.00/ebiten.DeviceScaleFactor(), float64(height)/2.00/ebiten.DeviceScaleFactor()

	g.Person = coordinates{
		x: int(fx),
		y: int(fy),
		// x: 240,
		// y: 240,
	}

	g.Font, err = loadFont("/ARCADECLASSIC.TTF")
	if err != nil {
		log.Panicln(err)
	}

	err = g.LoadImages()
	if err != nil {
		log.Panicln(err)
	}

	g.GamMap = initGameMap()

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Hello, World!")
	if err = ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Draw(screen *ebiten.Image) {
	// screen.Fill(color.White)
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	// text.Draw(screen, "P", g.Font, g.Person.x, g.Person.y, colornames.Blanchedalmond)
	g.drawMap(screen)
	tilegen.Person(screen, g.Person.x, g.Person.y)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	var g Game
	g.Init()
}

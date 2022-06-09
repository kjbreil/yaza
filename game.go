package main

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kjbreil/yaza/assets"
	"github.com/rakyll/statik/fs"
	"golang.org/x/image/font"
	"io/ioutil"
)

type Game struct {
	Font font.Face

	Person coordinates

	GamMap *gameMap

	AD     assets.Assets
	Assets []*ebiten.Image
}

type coordinates struct {
	x int
	y int
}

func loadFont(name string) (font.Face, error) {
	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	f, err := statikFS.Open(name)
	if err != nil {
		return nil, fmt.Errorf("error loading ttf: %v \n", err)

	}
	face, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	parsedFont, err := freetype.ParseFont(face)
	if err != nil {
		return nil, err
	}

	faceFont := truetype.NewFace(parsedFont, &truetype.Options{
		Size: 16,
		DPI:  96,
	})
	return faceFont, nil
}

package main

import (
	"fmt"
	"github.com/kjbreil/yaza/assets"
	"github.com/rakyll/statik/fs"
)

func (g *Game) LoadImages() error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	g.Assets, err = assets.SP.Load()
	if err != nil {
		return err
	}

	f, err := statikFS.Open("/assets.json")
	if err != nil {
		fmt.Println("error wit statik: ", err)
	}
	g.AD = assets.Read(f)
	return nil
}

package assets

import (
	"encoding/json"
	"image"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Assets is a map of the named assets
type Assets map[string]*Asset

// Asset is information about a name asset
type Asset struct {
	Sprite       string `json:"sprite,omitempty"`
	Index        int64  `json:"index,omitempty"`
	IsRow        *bool  `json:"is_row,omitempty"`
	CanDirection *bool  `json:"can_direction,omitempty"`
	Speed        *int64 `json:"speed,omitempty"`
}

// Row returns a random asset from a passed row
func (a *Asset) Row() int64 {
	if a.IsRow != nil && *a.IsRow {
		rand.Seed(time.Now().UnixNano())
		rowStartIndex := a.Index - (a.Index % 16)
		return rowStartIndex + rand.Int63n(16)
	}
	return a.Index
}

// Get will return an asset for a given reference, giving a row if needed
func (a *Asset) Get() *Asset {
	switch {
	case a.IsRow != nil && *a.IsRow:
		return &Asset{
			Sprite: a.Sprite,
			Index:  a.Row(),
			IsRow:  a.IsRow,
		}
	}
	return a
}

// Write to a file with the json of the assets information
func (as *Assets) Write(filename string) {
	d, err := json.MarshalIndent(as, "", "    ")
	if err != nil {
		log.Panicf("error marsheling json: %v", err)
	}

	err = ioutil.WriteFile(filename, d, 0666)
	if err != nil {
		log.Fatalf("error writing file %s: %v", filename, err)
	}
}

// Read from a file of asset information
func Read(f http.File) Assets {
	var as Assets
	d, err := ioutil.ReadAll(f)
	if err != nil {
		log.Panicf("error reading file: %v", err)
	}
	err = json.Unmarshal(d, &as)
	if err != nil {
		log.Panicf("error unmarshaling file: %v", err)
	}
	return as
}

// AllInSprite lists all the named assets in a sprite file
func (as Assets) AllInSprite(sprite string) (names []string) {

	for n, ea := range as {
		if ea.Sprite == sprite {
			names = append(names, n)
		}
	}
	return
}

// Direction returns a asset in the direction (next 4 are N,E,S,W or 0,1,2,3)
func (a Asset) Direction(d int16) *Asset {
	// a := as[name]
	ass := Asset{
		Sprite: a.Sprite,
		Index:  a.Index + int64(d+1),
	}
	return &ass
}

// Rect returns the assett and rectangle for said asset
func (as Assets) Rect(name string) (*Asset, *image.Rectangle) {
	indexMax := 16
	pixelSize := 16

	for n, ea := range as {
		if n == name {
			nas := ea.Get()
			index := int(nas.Index)
			x := (index % indexMax) * indexMax
			y := (index / indexMax) * indexMax
			rect := image.Rect(x, y, x+pixelSize, y+pixelSize)
			return nas, &rect
		}
	}
	return nil, nil
}

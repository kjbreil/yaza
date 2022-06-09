package assets

import (
	"fmt"
	"image"
	_ "image/png"
	"net/http"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/kjbreil/yaza/resources/statik"
	"github.com/rakyll/statik/fs"
)

// Load a sprite list into the holder
func (sp sprites) Load() ([]*ebiten.Image, error) {
	statikFS, err := fs.New()
	if err != nil {
		return nil, fmt.Errorf("new fs failed: %v", err)
	}

	var imgs []*ebiten.Image
	for i := range sp {
		img, err := decodeImage(statikFS, string(sp[i]))
		if err != nil {
			// fmt.Printf("decoding image %s had error: %v", Sprites[i], err)
			return nil, fmt.Errorf("decoding image %s had error: %v", sp[i], err)
		}
		imgs = append(imgs, img)
	}

	return imgs, nil
}

func decodeImage(statikFS http.FileSystem, filename string) (*ebiten.Image, error) {

	// f, err := statikFS.Open(filename)
	// TODO: Test and figure out why fp isn't just filename???
	f, err := statikFS.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("error opening file: %s %v", filename, err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	eimg := ebiten.NewImageFromImage(img)
	if err != nil {
		return nil, err
	}
	return eimg, nil
}

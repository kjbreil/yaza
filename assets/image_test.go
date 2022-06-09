package assets

import (
	"fmt"
	"image"
	"testing"
)

func Test_decodeImage(t *testing.T) {

	imgs, _ := SP.Load()

	si := imgs[0]

	ii := image.Image(si)

	fmt.Println(ii)
}

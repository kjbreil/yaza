package assets

import (
	"testing"

	_ "github.com/kjbreil/yaza/resources/statik"
)

// TODO: make array for words to test
func Test_sprites_I(t *testing.T) {
	i := SP.I("ground")
	if i != 0 {
		t.Fatalf("found index %d is not 0 for building", i)
	}
}

package assets

import (
	"log"
	"strings"
)

type sprites []sprite

type sprite string

// SP is an ordered array of the sprites
var SP = sprites{
	"/ground.png",
	"/building.png",
	"/people.png",
	"/util.png",
	"/menu.png",
	"/fixture.png",
}

// I returns the i of the named sprite
func (sp sprites) I(name string) int {
	for i := range sp {
		if strings.Contains(sp[i].name(), name) {
			return i
		}
	}
	log.Panicf("could not find named sprite: %s", name)

	return 0
}

// name returns the name of the sprite from filename
// remove the first 2 characters and using that remove the last 4
func (s *sprite) name() string {
	st := *s
	st = st[1:][:len(st[1:])-4]
	return string(st)
}

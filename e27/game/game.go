package game

import (
	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

const windowTitle = "A walk in the park"

var (
	goalAckMarker = &awakengine.Transient{
		Anim: &Anim{
			Key:       "mark",
			Offset:    vec.I2{15, 15},
			Frames:    4,
			FrameSize: vec.I2{32, 32},
			Mode:      awakengine.AnimOneShot,
		},
		birth: -999,
	}
)

// Game implements awakengine.Game
type Game struct {
	pixelSize    int
	camSize      vec.I2
	levelPreview bool
}

// New creates a new Game.
func New(levelPreview bool) *Game {
	ps := 3
	cs := vec.I2{267, 150}
	if levelPreview {
		ps = 1
		cs = vec.I2{1024, 1024}
	}
	return &Game{
		pixelSize:    ps,
		camSize:      cs,
		levelPreview: levelPreview,
	}
}

// Units provides all units in the level.
func (*Game) Units() []awakengine.Unit {
	return nil
}

// Viewport is the size of the window and the pixels in the window.
func (*Game) Viewport() (camSize vec.I2, pixelSize int) {
	return camSize, pixelSize
}

package game

import (
	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

const (
	// One frame of animation for every (animationPeriod) frames rendered.
	// So animation FPS = 60 / animationPeriod.
	animPeriod = 3

	windowTitle = "A walk in the park"
)

var (
	goalAckMarker = &awakengine.Transient{
		A: &awakengine.Anim{
			Key:       "mark",
			Offset:    vec.I2{15, 15},
			Frames:    4,
			FrameSize: vec.I2{32, 32},
			Mode:      awakengine.AnimOneShot,
		},
		Birth: -999,
	}

	theW = &awakengine.Doodad{
		BaseDoodad: baseDoodads["W"],
		P:          vec.I2{860, 453},
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

// Player returns the player unit.
func (*Game) Player() awakengine.Unit {
	return player
}

// Sprites provides all sprites in the level.
func (*Game) Sprites() []awakengine.Sprite {
	return []awakengine.Sprite{
		player,
		goalAckMarker,
		theW,
	}
}

// Viewport is the size of the window and the pixels in the window.
func (g *Game) Viewport() (cs vec.I2, ps, ap int, title string) {
	return g.camSize, g.pixelSize, animPeriod, windowTitle
}

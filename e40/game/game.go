// Copyright 2016 Josh Deprez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package game

import (
	"github.com/DrJosh9000/awakeman/common"
	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

const windowTitle = "Awakeman! #40: Escape from the Dark Library"

// Game implements awakengine.Game
type Game struct {
	pixelSize    int
	levelPreview bool
	scene        *awakengine.Scene
}

// New creates a new Game.
func New(levelPreview bool) *Game {
	ps := 3
	cs := vec.I2{267, 150}
	if levelPreview {
		ps = 2
		cs = vec.I2{1024, 512}
		//cs = vec.I2{512, 256}
	}
	scene := &awakengine.Scene{
		CameraSize: cs,
	}
	/*for _, d := range l.Doodads {
		scene.AddObject(d)
	}*/
	scene.AddObject(
		&awakengine.SpriteObject{Sprite: player, Semiobject: player},
		&awakengine.SpriteObject{Sprite: goalAckMarker, Semiobject: goalAckMarker},
	)
	hud := &awakengine.HUDRegion{
		Bubble: &awakengine.Bubble{
			UL:  vec.I2{2, 2},
			DR:  vec.I2{61, 37},
			Key: "inv_bubble",
		},
		Items: []awakengine.Drawable{
			&awakengine.Billboard{
				SheetFrame: itemPhone,
				P:          vec.I2{8, 8},
			},
			&awakengine.Billboard{
				SheetFrame: itemDucky,
				P:          vec.I2{8 + 24, 8},
			},
		},
		V: true,
		R: false,
	}
	hud.AddToScene(scene)
	return &Game{
		pixelSize:    ps,
		levelPreview: levelPreview,
		scene:        scene,
	}
}

func (*Game) BubbleKey() (string, string) { return "inv_bubble", "bubble" }

// Font returns the default typeface.
func (*Game) Font() awakengine.Font {
	return common.MunroFont{
		Invert: true,
	}
}

func (*Game) Handle(e awakengine.Event) {
	if e.Type == awakengine.EventMouseUp {
		goalAckMarker.Begin(e.Pos, e.Time)
		player.SetPath(awakengine.Navigate(player.Pos(), e.Pos))
	}
	if len(player.Path()) == 0 {
		goalAckMarker.End()
	}
}

// Player returns the player unit.
func (*Game) Player() awakengine.Unit {
	return player
}

func (g *Game) Scene() *awakengine.Scene { return g.scene }

// Viewport is the size of the window and the pixels in the window.
func (g *Game) Viewport() (pixSize int, title string) {
	return g.pixelSize, windowTitle
}

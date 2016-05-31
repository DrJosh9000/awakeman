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

const windowTitle = "A walk in the park"

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

func (*Game) BubbleKey() (string, string) { return "bubble", "inv_bubble" }

// Font returns the default typeface.
func (*Game) Font() awakengine.Font {
	return common.MunroFont{}
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

// Objects provides all sprites in the level.
func (*Game) Objects() []awakengine.Object {
	return []awakengine.Object{
		&awakengine.SpriteObject{Sprite: player, Semiobject: player},
		&awakengine.SpriteObject{Sprite: goalAckMarker, Semiobject: goalAckMarker},
		&awakengine.SpriteObject{Sprite: theW, Semiobject: theW},
	}
}

// Player returns the player unit.
func (*Game) Player() awakengine.Unit {
	return player
}

// Viewport is the size of the window and the pixels in the window.
func (g *Game) Viewport() (cs vec.I2, ps int, title string) {
	return g.camSize, g.pixelSize, windowTitle
}

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

var gameOver = false

// Game implements awakengine.Game
type Game struct {
	pixelSize    int
	levelPreview bool
	noTriggers   bool
	scene        *awakengine.Scene
}

var scene *awakengine.Scene

// New creates a new Game.
func New(levelPreview, noTriggers bool) *Game {
	ps := 3
	cs := vec.I2{267, 150}
	ts := vec.I2{1024, 512}
	if levelPreview {
		ps = 2
		cs = ts
	}
	scene = awakengine.NewScene(cs, ts)
	/*for _, d := range l.Doodads {
		scene.AddObject(d)
	}*/
	player.View = scene.World
	playerDelegate.AckMarker.View.SetParent(scene.World)
	scene.CameraFocus(player.Pos.I2())
	scene.AddPart(player, playerDelegate.AckMarker)

	inventory.bubble = &awakengine.Bubble{
		Key:  "inv_bubble",
		View: &awakengine.View{},
	}
	inventory.bubble.SetParent(scene.HUD)
	inventory.bubble.AddToScene(scene)
	inventory.grid = &awakengine.Grid{
		GridDelegate: inventory,
		View:         &awakengine.View{},
	}
	inventory.grid.SetParent(inventory.bubble.View)

	inventory.AddItems(&ItemPhone{}, ItemDucky{})

	return &Game{
		pixelSize:    ps,
		levelPreview: levelPreview,
		noTriggers:   noTriggers,
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

func (g *Game) Handle(e *awakengine.Event) bool {
	if gameOver {
		return true
	}
	if e.Type == awakengine.EventMouseUp {
		if inventory.grid.Handle(e) {
			return true
		}
		playerDelegate.SetPath(awakengine.Navigate(player.Pos.I2(), e.WorldPos))
	}
	g.scene.CameraFocus(player.Pos.I2())
	return true
}

// Player returns the player unit.
func (*Game) Player() (awakengine.Unit, *awakengine.Sprite) {
	return playerDelegate, player
}

func (g *Game) Scene() *awakengine.Scene { return g.scene }

// Viewport is the size of the window and the pixels in the window.
func (g *Game) Viewport() (pixSize int, title string) {
	return g.pixelSize, windowTitle
}

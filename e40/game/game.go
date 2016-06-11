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
	"log"

	"github.com/DrJosh9000/awakeman/common"
	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

const windowTitle = "Awakeman! #40: Escape from the Dark Library"

// Game implements awakengine.Game
type Game struct {
	pixelSize    int
	levelPreview bool
	noTriggers   bool
	scene        *awakengine.Scene
}

// New creates a new Game.
func New(levelPreview, noTriggers bool) *Game {
	ps := 3
	cs := vec.I2{267, 150}
	ts := vec.I2{1024, 512}
	if levelPreview {
		ps = 2
		cs = ts
		//cs = vec.I2{512, 256}
	}
	scene := awakengine.NewScene(cs, ts)
	/*for _, d := range l.Doodads {
		scene.AddObject(d)
	}*/
	player.View = scene.World
	scene.CameraFocus(player.Pos.I2())
	goalAckMarker.View.SetParent(scene.World)
	scene.AddPart(player, goalAckMarker)
	/*
		hud := &awakengine.HUDRegion{
			Bubble: &awakengine.Bubble{
				Key: "inv_bubble",
				UL:  vec.I2{1, 7},
				DR:  vec.I2{25, 56},
			},
			V: true,
			R: false,
		}
		hud.AddToScene(scene)

		itemsGrid := &awakengine.Grid{
			GridDelegate: inventory,
			ChildOf:      awakengine.ChildOf{hud.Bubble},
		}
		itemsGrid.AddToScene(scene)
	*/

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

func (g *Game) Handle(e awakengine.Event) {
	if e.Type == awakengine.EventMouseUp {
		//goalAckMarker.Begin(e.ScenePos, e.Time)
		log.Printf("e.WorldPos: %v", e.WorldPos)
		goalAckMarker.ResetAnim()
		goalAckMarker.Pos = e.WorldPos.F2()
		goalAckMarker.SetVisible(true)
		playerDelegate.SetPath(awakengine.Navigate(player.Pos.I2(), e.WorldPos))
	}
	if len(playerDelegate.Path()) == 0 {
		goalAckMarker.SetVisible(false)
		// log.Printf("goalAckMarker sheet: %#v", goalAckMarker.SpriteDelegate.SpriteSheet(goalAckMarker))
	}
	g.scene.CameraFocus(player.Pos.I2())
}

// Player returns the player unit.
func (*Game) Player() awakengine.Unit {
	return playerDelegate
}

func (g *Game) Scene() *awakengine.Scene { return g.scene }

// Viewport is the size of the window and the pixels in the window.
func (g *Game) Viewport() (pixSize int, title string) {
	return g.pixelSize, windowTitle
}

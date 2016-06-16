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
	}
	scene := awakengine.NewScene(cs, ts)
	/*for _, d := range l.Doodads {
		scene.AddObject(d)
	}*/
	player.View = scene.World
	scene.CameraFocus(player.Pos.I2())
	goalAckMarker.View.SetParent(scene.World)
	scene.AddPart(player, goalAckMarker)

	inventory.scene = scene
	itemsBubble := &awakengine.Bubble{
		Key:  "inv_bubble",
		View: &awakengine.View{},
	}
	itemsBubble.SetParent(scene.HUD)
	itemsBubble.AddToScene(scene)
	itemsGrid := &awakengine.Grid{
		GridDelegate: inventory,
		View:         &awakengine.View{},
	}
	itemsGrid.SetPosition(vec.I2{1, 7})
	itemsGrid.SetParent(itemsBubble.View)
	itemsGrid.Reload()
	//log.Printf("itemsGrid bounds after Reload: %#v", itemsGrid.Bounds())
	itemsBubble.Surround(itemsGrid.LogicalBounds())
	itemsGrid.SetPosition(vec.I2{5, 5}) // bubblePartSize
	log.Printf("itemsGrid: %#v", itemsGrid)
	//log.Printf("itemsGrid bounds after Surround: %#v", itemsGrid.Bounds())
	//log.Printf("itemsBubble bounds after Surround: %#v", itemsBubble.Bounds())

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
		goalAckMarker.ResetAnim()
		goalAckMarker.Pos = e.WorldPos.F2()
		goalAckMarker.SetVisible(true)
		playerDelegate.SetPath(awakengine.Navigate(player.Pos.I2(), e.WorldPos))
	}
	if len(playerDelegate.Path()) == 0 {
		goalAckMarker.SetVisible(false)
	}
	g.scene.CameraFocus(player.Pos.I2())
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

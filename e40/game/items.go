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
	"fmt"
	"math/rand"

	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

var (
	itemsSheet = &awakengine.Sheet{
		Key:        "items",
		FrameSize:  vec.I2{24, 24},
		FrameInfos: awakengine.BasicFrameInfos(4, -1, vec.I2{0, 0}),
	}

	inventory = &Inventory{limit: 4}
)

const (
	itemPhone = iota
	itemKey
	itemBook
	itemDucky
)

type Item interface {
	Icon() int
	Handle(e *awakengine.Event) bool
}

var (
	batteryNouns = []string{
		"battery", "energy storage", "electricity source", "lithium-ion component", "ion exchanger",
	}
	batteryAdjectives = []string{
		"dead", "kaput", "drained", "exhausted", "all used up", "fully consumed", "in need of recharging",
	}

	duckyVerbs = []string{
		"Actify", "Activate", "Actuate", "Energise", "Launch", "Mobilise", "Start", "Switch On", "Trigger",
	}
)

func randWord(words []string) string { return words[rand.Intn(len(words))] }

// Some item implementations...

type ItemPhone struct {
	activations int
}

func (i *ItemPhone) Icon() int { return itemPhone }
func (i *ItemPhone) Handle(*awakengine.Event) bool {
	awakengine.PushDialogue([]*awakengine.DialogueLine{
		{Avatar: avatarAwakeman, Text: fmt.Sprintf("This is my phone. The %s is %s.", randWord(batteryNouns), randWord(batteryAdjectives))},
		{Text: "(Do something with the phone?)", Buttons: []*awakengine.ButtonSpec{
			{Label: "Do nothing"},
			{Label: "Use", Action: func() {
				switch i.activations {
				case 0:
					awakengine.PushDialogue(&awakengine.DialogueLine{
						Text: "(The phone starts booting, then stops and displays instructions for connecting a charge cable. After a few seconds, it turns off again.)",
					})
				case 1:
					awakengine.PushDialogue(&awakengine.DialogueLine{
						Text: "(Instructions for connecting a charge cable breifly flicker across the screen, then it is off once again.)",
					})
				case 2, 3, 4:
					awakengine.PushDialogue(&awakengine.DialogueLine{
						Text: "(The phone is inoperable.)",
					})
				case 100:
					awakengine.PushDialogue(&awakengine.DialogueLine{
						Text: "(Look, I didn't implement a phone OS inside this silly game, so trying to use it is pointless.)",
					})
				default:
					awakengine.PushDialogue(&awakengine.DialogueLine{
						Avatar: avatarAwakeman,
						Text:   "Let's try doing something else.",
					})
				}
				i.activations++
			}},
			{Label: "Discard", Action: func() {
				awakengine.PushDialogue(&awakengine.DialogueLine{
					Avatar: avatarAwakeman,
					Text:   "Hey, I'm not discarding my phone!",
				})
			}},
		}},
	}...)
	return true
}

type ItemDucky struct{}

func (ItemDucky) Icon() int { return itemDucky }
func (ItemDucky) Handle(*awakengine.Event) bool {
	awakengine.PushDialogue([]*awakengine.DialogueLine{
		{Avatar: avatarDucky, Text: "Quack!"},
		{Text: "(Do something with the Ducky?)", Buttons: []*awakengine.ButtonSpec{
			{Label: "Do nothing"},
			{Label: randWord(duckyVerbs), Action: func() {
				awakengine.PushDialogue(&awakengine.DialogueLine{Avatar: avatarDucky, Text: "Quack?"})
			}},
			{Label: "Discard", Action: func() {
				awakengine.PushDialogue(&awakengine.DialogueLine{
					Text: "(Awakeman removes it, but it somehow reappears in the same inventory slot straight away.)",
				})
			}},
		}},
	}...)
	return true
}

type ItemBook book

func (i *ItemBook) Icon() int { return itemBook }
func (i *ItemBook) Handle(e *awakengine.Event) bool {
	awakengine.PushDialogue([]*awakengine.DialogueLine{
		{Text: fmt.Sprintf(`(It's "%s".)`, i.title), Buttons: []*awakengine.ButtonSpec{
			{Label: "Do nothing"},
			{Label: "Re-read", Action: (*book)(i).readWithoutCommentary},
			{Label: "Discard", Action: func() {
				inventory.RemoveItem(i)
			}},
		}},
	}...)
	return true
}

type ItemKey struct {
	handleTimes int
}

func (i *ItemKey) Icon() int { return itemKey }
func (i *ItemKey) Handle(e *awakengine.Event) bool {
	if i.handleTimes == 0 {
		awakengine.PushDialogue([]*awakengine.DialogueLine{
			{Avatar: avatarAwakeman, Text: `An old-fashioned key.`},
			{Avatar: avatarAwakeman, Text: `Right about now, I wish that sentence stopped before the "key".`},
		}...)
	}
	i.handleTimes++
	awakengine.PushDialogueToBack(&awakengine.DialogueLine{
		Text: "(Do something with the key?)",
		Buttons: []*awakengine.ButtonSpec{
			{Label: "Do nothing"},
			{Label: "Use", Action: func() {
				if player.Pos.I2().Div(tileSize) != vec.NewI2(51, 3) {
					awakengine.PushDialogue(&awakengine.DialogueLine{
						Avatar: avatarAwakeman,
						Text:   `There are no locks within reach.`,
					})
					return
				}
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Text: `(Awakeman tries the key.)`},
					{Avatar: avatarAwakeman, Text: `Doesn't fit in the lock.`},
					{Avatar: avatarAwakeman, Text: `Doesn't fit anywhere.`},
					{Avatar: avatarAwakeman, Text: `Ffffnnnaargh!`},
					{Text: `(Awakeman punches the wall, brushing the key past a fob-swipe with a small red light.)`},
					{Avatar: avatarDoor, Text: `*beep* *beep* *beep*`},
					{Avatar: avatarDoor, Text: ``},
					{Avatar: avatarDoor, Text: `*click*`},
					{Avatar: avatarAwakeman, Text: `*blink*`},
				}...)
				doorUnlocked = true
			}},
		},
	})
	return true
}

type Inventory struct {
	bubble *awakengine.Bubble
	grid   *awakengine.Grid
	items  []Item
	limit  int
}

func (in *Inventory) Layout() {
	in.grid.SetPosition(vec.I2{1, 7})
	in.grid.Reload()
	in.bubble.Surround(inventory.grid.RelativeBounds())
	in.grid.SetPosition(vec.I2{5, 5}) // bubblePartSize
}

// AddItems does not check that the limit is exceeded.
func (in *Inventory) AddItems(items ...Item) {
	in.items = append(in.items, items...)
	in.Layout()
}

func (in *Inventory) RemoveItem(item Item) {
	ni := in.items[:0]
	for _, x := range in.items {
		if x == item {
			continue
		}
		ni = append(ni, x)
	}
	in.items = ni
	in.Layout()
}

// The below implement awakengine.GridDelegate.
func (in *Inventory) ItemSize() vec.I2 { return itemsSheet.FrameSize }
func (in *Inventory) Columns() int     { return 1 }
func (in *Inventory) NumItems() int    { return len(in.items) }

func (in *Inventory) Item(i int, par *awakengine.View) {
	iv := &awakengine.ImageView{
		SheetFrame: &awakengine.SheetFrame{
			Sheet:   itemsSheet,
			FrameNo: in.items[i].Icon(),
		},
		View: par,
	}
	scene.AddPart(iv)
}

func (in *Inventory) ItemHandle(i int, e *awakengine.Event) bool {
	return in.items[i].Handle(e)
}

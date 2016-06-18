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

	inventory = &Inventory{}
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
		"battery", "energy", "electricity source", "lithium-ion component", "ion exchanger",
	}
	batteryAdjectives = []string{
		"flat", "dead", "kaput", "drained", "exhausted", "all used up", "fully consumed", "in need of recharging",
	}
)

func randWord(words []string) string { return words[rand.Intn(len(words))] }

// Some item implementations...

type ItemPhone struct{}

func (i ItemPhone) Icon() int { return itemPhone }
func (i ItemPhone) Handle(*awakengine.Event) bool {
	awakengine.PushDialogue([]*awakengine.DialogueLine{
		{Avatar: avatarAwakeman, Text: fmt.Sprintf("This is my phone. The %s is %s.", randWord(batteryNouns), randWord(batteryAdjectives))},
		{Text: "(Discard the phone?)", Buttons: []*awakengine.ButtonSpec{
			{Label: "Keep it"},
			{Label: "Discard it", Action: func() {
				awakengine.PushDialogue(&awakengine.DialogueLine{
					Avatar: avatarAwakeman,
					Text:   "Hey, I'm not discarding that!",
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
		{Text: "(Discard the Ducky?)", Buttons: []*awakengine.ButtonSpec{
			{Label: "Keep it"},
			{Label: "Discard it", Action: func() {
				awakengine.PushDialogue(&awakengine.DialogueLine{
					Text: "(Awakeman tries to remove it, but it reappears in the inventory slot straight away.)",
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
		{Text: fmt.Sprintf(`(It's a copy of "%s" by %s.)`, i.title, i.author)},
		{Text: "(Discard the book?)", Buttons: []*awakengine.ButtonSpec{
			{Label: "Keep it"},
			{Label: "Discard it", Action: func() {
				inventory.RemoveItem(i)
			}},
		}},
	}...)
	return true
}

type Inventory struct {
	scene  *awakengine.Scene
	bubble *awakengine.Bubble
	grid   *awakengine.Grid
	items  []Item
}

func (in *Inventory) Layout() {
	in.grid.SetPosition(vec.I2{1, 7})
	in.grid.Reload()
	in.bubble.Surround(inventory.grid.RelativeBounds())
	in.grid.SetPosition(vec.I2{5, 5}) // bubblePartSize
}

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
	in.scene.AddPart(iv)
}

func (in *Inventory) ItemHandle(i int, e *awakengine.Event) bool {
	return in.items[i].Handle(e)
}

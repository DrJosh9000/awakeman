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
	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

var (
	avatarsSheet = &awakengine.Sheet{
		Key:        "inv_avatars",
		FrameInfos: awakengine.BasicFrameInfos(5, -1, vec.I2{0, 0}),
		FrameSize:  vec.I2{34, 64},
	}

	avatarAwakeman        = &awakengine.SheetFrame{avatarsSheet, 0}
	avatarSJ              = &awakengine.SheetFrame{avatarsSheet, 1}
	avatarDucky           = &awakengine.SheetFrame{avatarsSheet, 2}
	avatarAlamore         = &awakengine.SheetFrame{avatarsSheet, 3}
	avatarAwakemanOnPhone = &awakengine.SheetFrame{avatarsSheet, 4}
)

func (g *Game) Triggers() map[string]*awakengine.Trigger {
	if g.levelPreview || g.noTriggers {
		return nil
	}
	return map[string]*awakengine.Trigger{
		"startGame": {
			Fire: func(int) {
				awakengine.PushDialogueBack([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: "Yes hello, this is dog!", Buttons: []awakengine.ButtonSpec{
						{Label: "Yes, dog", Action: func() {
							awakengine.PushDialogueFront(
								[]*awakengine.DialogueLine{{Avatar: avatarSJ, Text: "Yes, dog!"}},
							)
						}},
						{Label: "No dog", Action: func() {
							awakengine.PushDialogueFront(
								[]*awakengine.DialogueLine{{Avatar: avatarSJ, Text: "No dog!"}},
							)
						}},
					}},
					{Text: "Awakeman! No. 40: Escape from the Dark Library\n\n(Click or tap to advance the dialogue.)", Slowness: -1},
					{Avatar: avatarAwakeman, Text: `*sigh*`, Slowness: 4},
					{Avatar: avatarDucky, Text: `Quack?`},
					{Avatar: avatarAwakeman, Text: `It's not you, Ducky. I'm glad my ear isn't being talked off by Alamore, but it's at the expense of my phone being dead.`},
					{Avatar: avatarDucky, Text: `Quaaaaack.`},
					{Avatar: avatarAwakeman, Text: `Yeah. It was my light source.`},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 4},
					{Avatar: avatarAwakeman, Text: `Hey, don't suppose you have a built-in light?`},
					{Avatar: avatarDucky, Text: `Quack quack.`},
					{Text: "(Click or tap on things to move & interact.)"},
				})
			},
		},

		"it's dark in here": {
			Active:  func(f int) bool { return f > 20*6 },
			Depends: []string{"startGame"},
			Fire: func(int) {
				awakengine.PushDialogueBack([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `My sight still hasn't adjusted to the dark, but I can make out things a few metres away.`},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 4},
					{Avatar: avatarAwakeman, Text: `Ducky, why are you here?`},
					{Avatar: avatarDucky, Text: `Quack!`},
					{Avatar: avatarAwakeman, Text: `Ask a silly question, get a silly answer.`},
					{Avatar: avatarDucky, Text: `*Ring, ring*`, Slowness: 2},
					{Avatar: avatarAwakeman, Text: `What?`},
					{Avatar: avatarSJ, Text: `Hello! Hello? Can you hear m--`, AutoNext: true},
					{Avatar: avatarAwakeman, Text: `Yeah, I can hear you. Science Jesus? You're talking to me through the duck?`},
					{Avatar: avatarSJ, Text: "Oh good, you picked up.\nYes, I'm talking through the duck!\nAsk a silly question, get a silly answer."},
					{Avatar: avatarSJ, Text: "Where are you?"},
					{Avatar: avatarAwakeman, Text: "I'm in the office, it's dark and cold and my phone is dead and", AutoNext: true},
					{Avatar: avatarSJ, Text: "That's not...huh...", AutoNext: true, Slowness: 2},
					{Avatar: avatarAwakeman, Text: "and there's books everywhere and the power is out and", AutoNext: true},
					{Avatar: avatarSJ, Text: "I need to check on something. But I think you can find my painting nearby!"},
					{Avatar: avatarAwakeman, Text: ". . .", Slowness: 8},
					{Avatar: avatarAwakeman, Text: "Your *painting*?", Slowness: 2},
					{Avatar: avatarSJ, Text: `You're still going to find it for me, right? Sorry, gotta make like a tree and "split". Ciao!`},
					{Text: "*click*"},
					{Avatar: avatarDucky, Text: `Quack!`},
					{Avatar: avatarAwakeman, Text: `It's "make like a tree and" . . .`},
					{Avatar: avatarAwakeman, Text: `*sigh*`},
					{Avatar: avatarAwakeman, Text: `Who cares.`},
				})
			},
		},

		"before the centre": {
			Active: func(int) bool { return player.Pos.I2().InRect(vec.I2{X: 675, Y: 396}, vec.I2{X: 691, Y: 405}) },
			Fire: func(int) {
				awakengine.PushDialogueBack([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `How long is this going to go on for?`},
				})
			},
		},

		"spiral centre": {
			Active: func(int) bool { return player.Pos.I2().InRect(vec.I2{745, 325}, vec.I2{775, 351}) },
			Fire: func(int) {
				awakengine.PushDialogueBack([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `What is the point of this spiral?`},
					{Avatar: avatarAwakeman, Text: `This has got to be the most terrible architecture it has ever been my misfortune to walk through.`},
					{Avatar: avatarAwakeman, Text: `Why would you make people walk almost the maximum possible distance?`},
					{Avatar: avatarAwakeman, Text: `What is the point?`},
					{Avatar: avatarAwakeman, Text: `I don't get it.`},
					{Avatar: avatarAwakeman, Text: `I . . .`},
					{Avatar: avatarAwakeman, Text: `I can't even`},
					{Avatar: avatarAwakeman, Text: `Why?`},
					{Avatar: avatarAwakeman, Text: `Is there some secret here?`},
					{Avatar: avatarAwakeman, Text: `Why would you do this?`},
					{Avatar: avatarAwakeman, Text: `*sob*`},
					{Avatar: avatarAwakeman, Text: `Why?`},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 8},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 4},
					{Avatar: avatarAwakeman, Text: `This isn't the office.`, Slowness: 4},
					{Avatar: avatarAwakeman, Text: `...is it.`},
				})
			},
		},

		"past the centre": {
			Active: func(int) bool { return player.Pos.I2().InRect(vec.I2{712, 261}, vec.I2{725, 271}) },
			Fire: func(int) {
				awakengine.PushDialogueBack([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: "If I'm not in the office, where am I?"},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 8},
					{Avatar: avatarAwakeman, Text: "Some kind of perverse library. Books and building."},
				})
			},
		},

		"past the centre 2": {
			Active: func(int) bool { return player.Pos.I2().InRect(vec.I2{X: 915, Y: 315}, vec.I2{X: 924, Y: 323}) },
			Fire: func(int) {
				awakengine.PushDialogueBack([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: "Perhaps..."},
					{Avatar: avatarAwakeman, Text: "Maybe I was teleported here, not the other way around!"},
					{Avatar: avatarDucky, Text: "Quack!"},
				})
			},
		},
	}
}

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
		Key:       "avatars",
		Frames:    5,
		FrameSize: vec.I2{34, 64},
	}
	timeOfW int
)

const (
	avatarNone     = iota - 1
	avatarAwakeman // 0
	avatarSJ
	avatarDucky
	avatarAlamore
	avatarAwakemanOnPhone
)

func (g *Game) Triggers() map[string]*awakengine.Trigger {
	if g.levelPreview {
		return nil
	}
	return map[string]*awakengine.Trigger{
		"startGame": {
			Active: func(int) bool { return true },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarNone, "Awakeman! No. 27: A walk in the park\n\n(Click or tap to advance the dialogue.)", nil},
				{avatarsSheet, avatarAwakeman, `Alright. It must just be fate that I got duckied while escaping from Science Jesus's never-ending explanations.`, nil},
				{avatarsSheet, avatarAwakeman, `May as well search for the missing painting of the letter "W"!`, nil},
				{avatarsSheet, avatarDucky, "Quack quack quack quack QUACK quack quack quack quack quack quack quack quack quack quack quack quack quack!", nil},
				{avatarsSheet, avatarAwakeman, `Erm.`, nil},
				{avatarsSheet, avatarAwakeman, `I don't speak "duck". So I'm guessing...`, nil},
				{avatarsSheet, avatarAwakeman, `...It must be...`, nil},
				{avatarsSheet, avatarAwakeman, `Nearby?`, nil},
				{avatarsSheet, avatarDucky, "Quack quack.", nil},
				{avatarsSheet, avatarNone, "(Click or tap on things to move & interact.)", nil},
			},
		},

		"building": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{208, 464}, vec.I2{304, 592}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "Is it in this building?", nil},
				{avatarsSheet, avatarDucky, "...", nil},
				{avatarsSheet, avatarAwakeman, `Guess that's a "no".`, nil},
			},
		},

		"otherBuilding": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{199, 747}, vec.I2{407, 892}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "Is it in this building?", nil},
				{avatarsSheet, avatarDucky, "...", nil},
				{avatarsSheet, avatarDucky, "...Qua?...", nil},
				{avatarsSheet, avatarAwakeman, "Hrmmm...", nil},
			},
		},

		"other other building": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{862, 38}, vec.I2{955, 233}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "I think this was an old army facility or something.", nil},
				{avatarsSheet, avatarAwakeman, "It hasn't been maintained.", nil},
			},
		},

		"after20seconds": {
			Active: func(gameFrame int) bool { return gameFrame > 20*60 },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "Ducky, can you at least give me a clue?", nil},
				{avatarsSheet, avatarDucky, "Quack quack quack quack.", nil},
				{avatarsSheet, avatarAwakeman, `Well that's not helpful.`, nil},
			},
		},

		"after22seconds": {
			Depends: []string{"after20seconds"},
			Active:  func(gameFrame int) bool { return gameFrame > 22*60 },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarNone, "*Ring*... *ring*...", nil},
				{avatarsSheet, avatarAwakeman, "Better answer that.", nil},
				{avatarsSheet, avatarAwakemanOnPhone, `Hello?`, nil},
				{avatarsSheet, avatarSJ, `I noticed you were wandering around the park. That's--`, nil},
				{avatarsSheet, avatarAwakemanOnPhone, `Hi Science Jesus. You're watching me?`, nil},
				{avatarsSheet, avatarSJ, `I'm always watching. I'm Science *Jesus*.`, nil},
				{avatarsSheet, avatarAwakemanOnPhone, `...`, nil},
				{avatarsSheet, avatarSJ, `But also watching from the ducky. I did tell you about that feature.`, nil},
				{avatarsSheet, avatarAwakemanOnPhone, `Duck-based surveillance. Right.`, nil},
				{avatarsSheet, avatarAwakemanOnPhone, `How did the ducky appear in front of me after I stormed out of your home?`, nil},
				{avatarsSheet, avatarSJ, `I explained that too, but you weren't listening.`, nil},
				{avatarsSheet, avatarAwakemanOnPhone, `Okay... Look, I'm going to--`, nil},
				{avatarsSheet, avatarSJ, `Find my painting? Great! Well, as you will have figured out by now, the ducky will help you with that.`, nil},
				{avatarsSheet, avatarAwakemanOnPhone, `Howw-wwwait, no, not asking.`, nil},
				{avatarsSheet, avatarSJ, `Hmm. But I wasn't done explaining. :(`, nil},
				{avatarsSheet, avatarAwakemanOnPhone, `Can you give me a clue?`, nil},
				{avatarsSheet, avatarSJ, `Not really. All I can tell from my sensor array here is you're actually pretty close!`, nil},
				{avatarsSheet, avatarSJ, `Just keep looking around the area.`, nil},
				{avatarsSheet, avatarAwakemanOnPhone, `Okay, SJ. Talk later.`, nil},
				{avatarsSheet, avatarNone, "*click*", nil},
			},
		},

		"park bench": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{480, 83}, vec.I2{512, 95}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, `"This park bench is dedicated to the memory of Dr W. D. Gaster."`, nil},
			},
		},

		"clearing": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{484, 423}, vec.I2{666, 607}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "Huh. What a cool little hideaway!", nil},
				{avatarsSheet, avatarDucky, "Quack!", nil},
				{avatarsSheet, avatarAwakeman, `Nice and peaceful. I could see myself having a forest nap after this "forest bath".`, nil},
				{avatarsSheet, avatarDucky, "Quack!", nil},
				{avatarsSheet, avatarAwakeman, "No Science Jesus to talk my ears off.", nil},
				{avatarsSheet, avatarDucky, "Quack quack QUACK quack, quack quack quack!", nil},
				{avatarsSheet, avatarAwakeman, "...", nil},
				{avatarsSheet, avatarAwakeman, "Yeah, it's too cold. Gotta stay awake.", nil},
			},
		},
		"weirdTree": {
			Depends: []string{"clearing"},
			Active:  func(int) bool { return player.Pos().InRect(vec.I2{528, 473}, vec.I2{639, 525}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "Weird tree!", nil},
			},
		},

		"clearing obstacle": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{613, 554}, vec.I2{667, 592}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "I don't think I can get past this fallen tree.", nil},
			},
		},

		// See the W? Let's hide it behind dialogue bubble trololololo >:)
		"was that a W": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{740, 390}, vec.I2{925, 444}) },
			Fire: func(gameFrame int) {
				theW.V = false
				timeOfW = gameFrame
			},
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, `Wait, is that...!`, nil},
				{avatarsSheet, avatarDucky, `QUACK! QUACK QUACK QUACK!`, nil},
				{avatarsSheet, avatarAwakeman, `Quick!`, nil},
			},
		},

		"it was?": {
			Depends: []string{"was that a W"},
			Active: func(gameFrame int) bool {
				return player.Pos().InRect(vec.I2{740, 400}, vec.I2{925, 444}) && gameFrame > timeOfW+120
			},
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, `It was...it...uh...`, nil},
				{avatarsSheet, avatarDucky, `QUaaaaaack.`, nil},
				{avatarsSheet, avatarAwakeman, `I could have sworn something was hanging on that tree.`, nil},
				{avatarsSheet, avatarNone, `*Ring*...`, nil},
				{avatarsSheet, avatarAwakeman, `Wh--phone? Not when we're so close!`, nil},
				{avatarsSheet, avatarNone, `*boop*`, nil},
			},
		},

		"this way?": {
			Depends: []string{"it was?"},
			Active:  func(int) bool { return player.Pos().InRect(vec.I2{740, 440}, vec.I2{925, 472}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, `Maybe it disappeared down this way.`, nil},
			},
		},

		"holey trunk": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{849, 624}, vec.I2{912, 667}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarDucky, `Quack!`, nil},
				{avatarsSheet, avatarAwakeman, `I wonder what made this hole in this old tree trunk. Something round?`, nil},
			},
		},

		"alamore": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{845, 853}, vec.I2{904, 898}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "Hey, Alamore!", nil},
				{avatarsSheet, avatarAlamore, "If it isn't Awakeman. Fancy seeing you out here! Enjoying the park?", nil},
				{avatarsSheet, avatarAwakeman, "It's been an interesting walk.", nil},
				{avatarsSheet, avatarAlamore, "What brings you out here?", nil},
				{avatarsSheet, avatarAwakeman, `Well, I'm looking for a missing painting of a letter "W"--`, nil},
				{avatarsSheet, avatarAlamore, "Surely a hero as brave and important as you has better things to do? :)", nil},
				{avatarsSheet, avatarAwakeman, `I'm not that important...`, nil},
				{avatarsSheet, avatarAwakeman, `Um. I mean..."Awakeman" *is just* my name, I'm not a hero at all.`, nil},
				{avatarsSheet, avatarAlamore, "No! Nonsense. You're smart, and very handsome, and a hero.", nil},
				{avatarsSheet, avatarAlamore, "It's like you're an alien from another planet, you don't fit in here, not with the rest of the small-minded people who live in this town. :)", nil},
				{avatarsSheet, avatarAwakeman, `Um. Thanks.`, nil},
				{avatarsSheet, avatarAlamore, "Who told you you're not a hero?", nil},
				{avatarsSheet, avatarAwakeman, "N--nobody?", nil},
				{avatarsSheet, avatarAlamore, "Oh come on, you can trust me. :)", nil},
				{avatarsSheet, avatarAwakeman, "No, I'm ... I don't think I'm special.", nil},
				{avatarsSheet, avatarAlamore, "Hm. Well, we should have a chat about a very important project...", nil},
				{avatarsSheet, avatarNone, "AWAKEMAN! No. 27\nCREDITS\nProgramming, art, story, characters, concept\nJosh Deprez", nil},
				{avatarsSheet, avatarNone, "TUNE IN NEXT WEEK for the next GRIPPING EPISODE of AWAKEMAN!", nil},
				{avatarsSheet, avatarNone, "(It might be the end of this episode, but I won't stop you playing some more.)", nil},
			},
		},
	}
}

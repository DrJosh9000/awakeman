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
				{avatarsSheet, avatarNone, "Awakeman! No. 27: A walk in the park\n\n(Click or tap to advance the dialogue.)"},
				{avatarsSheet, avatarAwakeman, `Alright. It must just be fate that I got duckied while escaping from Science Jesus's never-ending explanations.`},
				{avatarsSheet, avatarAwakeman, `May as well search for the missing painting of the letter "W"!`},
				{avatarsSheet, avatarDucky, "Quack quack quack quack QUACK quack quack quack quack quack quack quack quack quack quack quack quack quack!"},
				{avatarsSheet, avatarAwakeman, `Erm.`},
				{avatarsSheet, avatarAwakeman, `I don't speak "duck". So I'm guessing...`},
				{avatarsSheet, avatarAwakeman, `...It must be...`},
				{avatarsSheet, avatarAwakeman, `Nearby?`},
				{avatarsSheet, avatarDucky, "Quack quack."},
				{avatarsSheet, avatarNone, "(Click or tap on things to move & interact.)"},
			},
		},

		"building": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{208, 464}, vec.I2{304, 592}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "Is it in this building?"},
				{avatarsSheet, avatarDucky, "..."},
				{avatarsSheet, avatarAwakeman, `Guess that's a "no".`},
			},
		},

		"otherBuilding": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{199, 747}, vec.I2{407, 892}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "Is it in this building?"},
				{avatarsSheet, avatarDucky, "..."},
				{avatarsSheet, avatarDucky, "...Qua?..."},
				{avatarsSheet, avatarAwakeman, "Hrmmm..."},
			},
		},

		"other other building": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{862, 38}, vec.I2{955, 233}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "I think this was an old army facility or something."},
				{avatarsSheet, avatarAwakeman, "It hasn't been maintained."},
			},
		},

		"after20seconds": {
			Active: func(gameFrame int) bool { return gameFrame > 20*60 },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "Ducky, can you at least give me a clue?"},
				{avatarsSheet, avatarDucky, "Quack quack quack quack."},
				{avatarsSheet, avatarAwakeman, `Well that's not helpful.`},
			},
		},

		"after22seconds": {
			Depends: []string{"after20seconds"},
			Active:  func(gameFrame int) bool { return gameFrame > 22*60 },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarNone, "*Ring*... *ring*..."},
				{avatarsSheet, avatarAwakeman, "Better answer that."},
				{avatarsSheet, avatarAwakemanOnPhone, `Hello?`},
				{avatarsSheet, avatarSJ, `I noticed you were wandering around the park. That's--`},
				{avatarsSheet, avatarAwakemanOnPhone, `Hi Science Jesus. You're watching me?`},
				{avatarsSheet, avatarSJ, `I'm always watching. I'm Science *Jesus*.`},
				{avatarsSheet, avatarAwakemanOnPhone, `...`},
				{avatarsSheet, avatarSJ, `But also watching from the ducky. I did tell you about that feature.`},
				{avatarsSheet, avatarAwakemanOnPhone, `Duck-based surveillance. Right.`},
				{avatarsSheet, avatarAwakemanOnPhone, `How did the ducky appear in front of me after I stormed out of your home?`},
				{avatarsSheet, avatarSJ, `I explained that too, but you weren't listening.`},
				{avatarsSheet, avatarAwakemanOnPhone, `Okay... Look, I'm going to--`},
				{avatarsSheet, avatarSJ, `Find my painting? Great! Well, as you will have figured out by now, the ducky will help you with that.`},
				{avatarsSheet, avatarAwakemanOnPhone, `Howw-wwwait, no, not asking.`},
				{avatarsSheet, avatarSJ, `Hmm. But I wasn't done explaining. :(`},
				{avatarsSheet, avatarAwakemanOnPhone, `Can you give me a clue?`},
				{avatarsSheet, avatarSJ, `Not really. All I can tell from my sensor array here is you're actually pretty close!`},
				{avatarsSheet, avatarSJ, `Just keep looking around the area.`},
				{avatarsSheet, avatarAwakemanOnPhone, `Okay, SJ. Talk later.`},
				{avatarsSheet, avatarNone, "*click*"},
			},
		},

		"park bench": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{480, 83}, vec.I2{512, 95}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, `"This park bench is dedicated to the memory of Dr W. D. Gaster."`},
			},
		},

		"clearing": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{484, 423}, vec.I2{666, 607}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "Huh. What a cool little hideaway!"},
				{avatarsSheet, avatarDucky, "Quack!"},
				{avatarsSheet, avatarAwakeman, `Nice and peaceful. I could see myself having a forest nap after this "forest bath".`},
				{avatarsSheet, avatarDucky, "Quack!"},
				{avatarsSheet, avatarAwakeman, "No Science Jesus to talk my ears off."},
				{avatarsSheet, avatarDucky, "Quack quack QUACK quack, quack quack quack!"},
				{avatarsSheet, avatarAwakeman, "..."},
				{avatarsSheet, avatarAwakeman, "Yeah, it's too cold. Gotta stay awake."},
			},
		},
		"weirdTree": {
			Depends: []string{"clearing"},
			Active:  func(int) bool { return player.Pos().InRect(vec.I2{528, 473}, vec.I2{639, 525}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "Weird tree!"},
			},
		},

		"clearing obstacle": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{613, 554}, vec.I2{667, 592}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "I don't think I can get past this fallen tree."},
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
				{avatarsSheet, avatarAwakeman, `Wait, is that...!`},
				{avatarsSheet, avatarDucky, `QUACK! QUACK QUACK QUACK!`},
				{avatarsSheet, avatarAwakeman, `Quick!`},
			},
		},

		"it was?": {
			Depends: []string{"was that a W"},
			Active: func(gameFrame int) bool {
				return player.Pos().InRect(vec.I2{740, 400}, vec.I2{925, 444}) && gameFrame > timeOfW+120
			},
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, `It was...it...uh...`},
				{avatarsSheet, avatarDucky, `QUaaaaaack.`},
				{avatarsSheet, avatarAwakeman, `I could have sworn something was hanging on that tree.`},
				{avatarsSheet, avatarNone, `*Ring*...`},
				{avatarsSheet, avatarAwakeman, `Wh--phone? Not when we're so close!`},
				{avatarsSheet, avatarNone, `*boop*`},
			},
		},

		"this way?": {
			Depends: []string{"it was?"},
			Active:  func(int) bool { return player.Pos().InRect(vec.I2{740, 440}, vec.I2{925, 472}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, `Maybe it disappeared down this way.`},
			},
		},

		"holey trunk": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{849, 624}, vec.I2{912, 667}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarDucky, `Quack!`},
				{avatarsSheet, avatarAwakeman, `I wonder what made this hole in this old tree trunk. Something round?`},
			},
		},

		"alamore": {
			// depends: []string{ "wasThatADoubleU" },
			Active: func(int) bool { return player.Pos().InRect(vec.I2{845, 853}, vec.I2{904, 898}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsSheet, avatarAwakeman, "Hey, Alamore!"},
				{avatarsSheet, avatarAlamore, "If it isn't Awakeman. Fancy seeing you out here! Enjoying the park?"},
				{avatarsSheet, avatarAwakeman, "It's been an interesting walk."},
				{avatarsSheet, avatarAlamore, "What brings you out here?"},
				{avatarsSheet, avatarAwakeman, `Well, I'm looking for a missing painting of a letter "W"--`},
				{avatarsSheet, avatarAlamore, "Surely a hero as brave and important as you has better things to do? :)"},
				{avatarsSheet, avatarAwakeman, `I'm not that important...`},
				{avatarsSheet, avatarAwakeman, `Um. I mean..."Awakeman" *is just* my name, I'm not a hero at all.`},
				{avatarsSheet, avatarAlamore, "No! Nonsense. You're smart, and very handsome, and a hero."},
				{avatarsSheet, avatarAlamore, "It's like you're an alien from another planet, you don't fit in here, not with the rest of the small-minded people who live in this town. :)"},
				{avatarsSheet, avatarAwakeman, `Um. Thanks.`},
				{avatarsSheet, avatarAlamore, "Who told you you're not a hero?"},
				{avatarsSheet, avatarAwakeman, "N--nobody?"},
				{avatarsSheet, avatarAlamore, "Oh come on, you can trust me. :)"},
				{avatarsSheet, avatarAwakeman, "No, I'm ... I don't think I'm special."},
				{avatarsSheet, avatarAlamore, "Hm. Well, we should have a chat about a very important project..."},
				{avatarsSheet, avatarNone, "AWAKEMAN! No. 27\nCREDITS\nProgramming, art, story, characters, concept\nJosh Deprez"},
				{avatarsSheet, avatarNone, "TUNE IN NEXT WEEK for the next GRIPPING EPISODE of AWAKEMAN!"},
				{avatarsSheet, avatarNone, "(It might be the end of this episode, but I won't stop you playing some more.)"},
			},
		},
	}
}

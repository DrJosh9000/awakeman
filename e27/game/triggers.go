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
	avatarsAnim = &awakengine.Anim{
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
				{avatarsAnim, avatarNone, "Awakeman! No. 27: A walk in the park\n\n(Click or tap to advance the dialogue.)"},
				{avatarsAnim, avatarAwakeman, `Alright. It must just be fate that I got duckied while escaping from Science Jesus's never-ending explanations.`},
				{avatarsAnim, avatarAwakeman, `May as well search for the missing painting of the letter "W"!`},
				{avatarsAnim, avatarDucky, "Quack quack quack quack QUACK quack quack quack quack quack quack quack quack quack quack quack quack quack!"},
				{avatarsAnim, avatarAwakeman, `Erm.`},
				{avatarsAnim, avatarAwakeman, `I don't speak "duck". So I'm guessing...`},
				{avatarsAnim, avatarAwakeman, `...It must be...`},
				{avatarsAnim, avatarAwakeman, `Nearby?`},
				{avatarsAnim, avatarDucky, "Quack quack."},
				{avatarsAnim, avatarNone, "(Click or tap on things to move & interact.)"},
			},
		},

		"building": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{208, 464}, vec.I2{304, 592}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarAwakeman, "Is it in this building?"},
				{avatarsAnim, avatarDucky, "..."},
				{avatarsAnim, avatarAwakeman, `Guess that's a "no".`},
			},
		},

		"otherBuilding": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{199, 747}, vec.I2{407, 892}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarAwakeman, "Is it in this building?"},
				{avatarsAnim, avatarDucky, "..."},
				{avatarsAnim, avatarDucky, "...Qua?..."},
				{avatarsAnim, avatarAwakeman, "Hrmmm..."},
			},
		},

		"other other building": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{862, 38}, vec.I2{955, 233}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarAwakeman, "I think this was an old army facility or something."},
				{avatarsAnim, avatarAwakeman, "It hasn't been maintained."},
			},
		},

		"after20seconds": {
			Active: func(gameFrame int) bool { return gameFrame > 20*60 },
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarAwakeman, "Ducky, can you at least give me a clue?"},
				{avatarsAnim, avatarDucky, "Quack quack quack quack."},
				{avatarsAnim, avatarAwakeman, `Well that's not helpful.`},
			},
		},

		"after22seconds": {
			Depends: []string{"after20seconds"},
			Active:  func(gameFrame int) bool { return gameFrame > 22*60 },
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarNone, "*Ring*... *ring*..."},
				{avatarsAnim, avatarAwakeman, "Better answer that."},
				{avatarsAnim, avatarAwakemanOnPhone, `Hello?`},
				{avatarsAnim, avatarSJ, `I noticed you were wandering around the park. That's--`},
				{avatarsAnim, avatarAwakemanOnPhone, `Hi Science Jesus. You're watching me?`},
				{avatarsAnim, avatarSJ, `I'm always watching. I'm Science *Jesus*.`},
				{avatarsAnim, avatarAwakemanOnPhone, `...`},
				{avatarsAnim, avatarSJ, `But also watching from the ducky. I did tell you about that feature.`},
				{avatarsAnim, avatarAwakemanOnPhone, `Duck-based surveillance. Right.`},
				{avatarsAnim, avatarAwakemanOnPhone, `How did the ducky appear in front of me after I stormed out of your home?`},
				{avatarsAnim, avatarSJ, `I explained that too, but you weren't listening.`},
				{avatarsAnim, avatarAwakemanOnPhone, `Okay... Look, I'm going to--`},
				{avatarsAnim, avatarSJ, `Find my painting? Great! Well, as you will have figured out by now, the ducky will help you with that.`},
				{avatarsAnim, avatarAwakemanOnPhone, `Howw-wwwait, no, not asking.`},
				{avatarsAnim, avatarSJ, `Hmm. But I wasn't done explaining. :(`},
				{avatarsAnim, avatarAwakemanOnPhone, `Can you give me a clue?`},
				{avatarsAnim, avatarSJ, `Not really. All I can tell from my sensor array here is you're actually pretty close!`},
				{avatarsAnim, avatarSJ, `Just keep looking around the area.`},
				{avatarsAnim, avatarAwakemanOnPhone, `Okay, SJ. Talk later.`},
				{avatarsAnim, avatarNone, "*click*"},
			},
		},

		"park bench": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{480, 83}, vec.I2{512, 95}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarAwakeman, `"This park bench is dedicated to the memory of Dr W. D. Gaster."`},
			},
		},

		"clearing": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{484, 423}, vec.I2{666, 607}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarAwakeman, "Huh. What a cool little hideaway!"},
				{avatarsAnim, avatarDucky, "Quack!"},
				{avatarsAnim, avatarAwakeman, `Nice and peaceful. I could see myself having a forest nap after this "forest bath".`},
				{avatarsAnim, avatarDucky, "Quack!"},
				{avatarsAnim, avatarAwakeman, "No Science Jesus to talk my ears off."},
				{avatarsAnim, avatarDucky, "Quack quack QUACK quack, quack quack quack!"},
				{avatarsAnim, avatarAwakeman, "..."},
				{avatarsAnim, avatarAwakeman, "Yeah, it's too cold. Gotta stay awake."},
			},
		},
		"weirdTree": {
			Depends: []string{"clearing"},
			Active:  func(int) bool { return player.Pos().InRect(vec.I2{528, 473}, vec.I2{639, 525}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarAwakeman, "Weird tree!"},
			},
		},

		"clearing obstacle": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{613, 554}, vec.I2{667, 592}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarAwakeman, "I don't think I can get past this fallen tree."},
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
				{avatarsAnim, avatarAwakeman, `Wait, is that...!`},
				{avatarsAnim, avatarDucky, `QUACK! QUACK QUACK QUACK!`},
				{avatarsAnim, avatarAwakeman, `Quick!`},
			},
		},

		"it was?": {
			Depends: []string{"was that a W"},
			Active: func(gameFrame int) bool {
				return player.Pos().InRect(vec.I2{740, 400}, vec.I2{925, 444}) && gameFrame > timeOfW+120
			},
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarAwakeman, `It was...it...uh...`},
				{avatarsAnim, avatarDucky, `QUaaaaaack.`},
				{avatarsAnim, avatarAwakeman, `I could have sworn something was hanging on that tree.`},
				{avatarsAnim, avatarNone, `*Ring*...`},
				{avatarsAnim, avatarAwakeman, `Wh--phone? Not when we're so close!`},
				{avatarsAnim, avatarNone, `*boop*`},
			},
		},

		"this way?": {
			Depends: []string{"it was?"},
			Active:  func(int) bool { return player.Pos().InRect(vec.I2{740, 440}, vec.I2{925, 472}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarAwakeman, `Maybe it disappeared down this way.`},
			},
		},

		"holey trunk": {
			Active: func(int) bool { return player.Pos().InRect(vec.I2{849, 624}, vec.I2{912, 667}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarDucky, `Quack!`},
				{avatarsAnim, avatarAwakeman, `I wonder what made this hole in this old tree trunk. Something round?`},
			},
		},

		"alamore": {
			// depends: []string{ "wasThatADoubleU" },
			Active: func(int) bool { return player.Pos().InRect(vec.I2{845, 853}, vec.I2{904, 898}) },
			Dialogues: []awakengine.DialogueLine{
				{avatarsAnim, avatarAwakeman, "Hey, Alamore!"},
				{avatarsAnim, avatarAlamore, "If it isn't Awakeman. Fancy seeing you out here! Enjoying the park?"},
				{avatarsAnim, avatarAwakeman, "It's been an interesting walk."},
				{avatarsAnim, avatarAlamore, "What brings you out here?"},
				{avatarsAnim, avatarAwakeman, `Well, I'm looking for a missing painting of a letter "W"--`},
				{avatarsAnim, avatarAlamore, "Surely a hero as brave and important as you has better things to do? :)"},
				{avatarsAnim, avatarAwakeman, `I'm not that important...`},
				{avatarsAnim, avatarAwakeman, `Um. I mean..."Awakeman" *is just* my name, I'm not a hero at all.`},
				{avatarsAnim, avatarAlamore, "No! Nonsense. You're smart, and very handsome, and a hero."},
				{avatarsAnim, avatarAlamore, "It's like you're an alien from another planet, you don't fit in here, not with the rest of the small-minded people who live in this town. :)"},
				{avatarsAnim, avatarAwakeman, `Um. Thanks.`},
				{avatarsAnim, avatarAlamore, "Who told you you're not a hero?"},
				{avatarsAnim, avatarAwakeman, "N--nobody?"},
				{avatarsAnim, avatarAlamore, "Oh come on, you can trust me. :)"},
				{avatarsAnim, avatarAwakeman, "No, I'm ... I don't think I'm special."},
				{avatarsAnim, avatarAlamore, "Hm. Well, we should have a chat about a very important project..."},
				{avatarsAnim, avatarNone, "AWAKEMAN! No. 27\nCREDITS\nProgramming, art, story, characters, concept\nJosh Deprez"},
				{avatarsAnim, avatarNone, "TUNE IN NEXT WEEK for the next GRIPPING EPISODE of AWAKEMAN!"},
				{avatarsAnim, avatarNone, "(It might be the end of this episode, but I won't stop you playing some more.)"},
			},
		},
	}
}

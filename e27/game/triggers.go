package game

import (
	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

var timeOfW int

var triggers = map[string]*awakengine.Trigger{
	"startGame": {
		Active: func() bool { return true },
		Dialogues: []awakengine.DialogueLine{
			{avatarNone, "Awakeman! No. 27: A walk in the park\n\n(Click or tap to advance the dialogue.)"},
			{avatarAwakeman, `Alright. It must just be fate that I got duckied while escaping from Science Jesus's never-ending explanations.`},
			{avatarAwakeman, `May as well search for the missing painting of the letter "W"!`},
			{avatarDucky, "Quack quack quack quack QUACK quack quack quack quack quack quack quack quack quack quack quack quack quack!"},
			{avatarAwakeman, `Erm.`},
			{avatarAwakeman, `I don't speak "duck". So I'm guessing...`},
			{avatarAwakeman, `...It must be...`},
			{avatarAwakeman, `Nearby?`},
			{avatarDucky, "Quack quack."},
			{avatarNone, "(Click or tap on things to move & interact.)"},
		},
	},

	"building": {
		Active: func() bool { return player.Pos().InRect(vec.I2{208, 464}, vec.I2{304, 592}) },
		Dialogues: []awakengine.DialogueLine{
			{avatarAwakeman, "Is it in this building?"},
			{avatarDucky, "..."},
			{avatarAwakeman, `Guess that's a "no".`},
		},
	},

	"otherBuilding": {
		Active: func() bool { return player.Pos().InRect(vec.I2{199, 747}, vec.I2{407, 892}) },
		Dialogues: []awakengine.DialogueLine{
			{avatarAwakeman, "Is it in this building?"},
			{avatarDucky, "..."},
			{avatarDucky, "...Qua?..."},
			{avatarAwakeman, "Hrmmm..."},
		},
	},

	"other other building": {
		Active: func() bool { return player.Pos().InRect(vec.I2{862, 38}, vec.I2{955, 233}) },
		Dialogues: []awakengine.DialogueLine{
			{avatarAwakeman, "I think this was an old army facility or something."},
			{avatarAwakeman, "It hasn't been maintained."},
		},
	},

	"after20seconds": {
		Active: func() bool { return gameFrame > 20*60 },
		Dialogues: []awakengine.DialogueLine{
			{avatarAwakeman, "Ducky, can you at least give me a clue?"},
			{avatarDucky, "Quack quack quack quack."},
			{avatarAwakeman, `Well that's not helpful.`},
		},
	},

	"after22seconds": {
		Depends: []string{"after20seconds"},
		Active:  func() bool { return gameFrame > 22*60 },
		Dialogues: []awakengine.DialogueLine{
			{avatarNone, "*Ring*... *ring*..."},
			{avatarAwakeman, "Better answer that."},
			{avatarAwakemanOnPhone, `Hello?`},
			{avatarSJ, `I noticed you were wandering around the park. That's--`},
			{avatarAwakemanOnPhone, `Hi Science Jesus. You're watching me?`},
			{avatarSJ, `I'm always watching. I'm Science *Jesus*.`},
			{avatarAwakemanOnPhone, `...`},
			{avatarSJ, `But also watching from the ducky. I did tell you about that feature.`},
			{avatarAwakemanOnPhone, `Duck-based surveillance. Right.`},
			{avatarAwakemanOnPhone, `How did the ducky appear in front of me after I stormed out of your home?`},
			{avatarSJ, `I explained that too, but you weren't listening.`},
			{avatarAwakemanOnPhone, `Okay... Look, I'm going to--`},
			{avatarSJ, `Find my painting? Great! Well, as you will have figured out by now, the ducky will help you with that.`},
			{avatarAwakemanOnPhone, `Howw-wwwait, no, not asking.`},
			{avatarSJ, `Hmm. But I wasn't done explaining. :(`},
			{avatarAwakemanOnPhone, `Can you give me a clue?`},
			{avatarSJ, `Not really. All I can tell from my sensor array here is you're actually pretty close!`},
			{avatarSJ, `Just keep looking around the area.`},
			{avatarAwakemanOnPhone, `Okay, SJ. Talk later.`},
			{avatarNone, "*click*"},
		},
	},

	"park bench": {
		Active: func() bool { return player.Pos().InRect(vec.I2{480, 83}, vec.I2{512, 95}) },
		Dialogues: []awakengine.DialogueLine{
			{avatarAwakeman, `"This park bench is dedicated to the memory of Dr W. D. Gaster."`},
		},
	},

	"clearing": {
		Active: func() bool { return player.Pos().InRect(vec.I2{484, 423}, vec.I2{666, 607}) },
		Dialogues: []awakengine.DialogueLine{
			{avatarAwakeman, "Huh. What a cool little hideaway!"},
			{avatarDucky, "Quack!"},
			{avatarAwakeman, `Nice and peaceful. I could see myself having a forest nap after this "forest bath".`},
			{avatarDucky, "Quack!"},
			{avatarAwakeman, "No Science Jesus to talk my ears off."},
			{avatarDucky, "Quack quack QUACK quack, quack quack quack!"},
			{avatarAwakeman, "..."},
			{avatarAwakeman, "Yeah, it's too cold. Gotta stay awake."},
		},
	},
	"weirdTree": {
		Depends: []string{"clearing"},
		Active:  func() bool { return player.Pos().InRect(vec.I2{528, 473}, vec.I2{639, 525}) },
		Dialogues: []awakengine.DialogueLine{
			{avatarAwakeman, "Weird tree!"},
		},
	},

	"clearing obstacle": {
		Active: func() bool { return player.Pos().InRect(vec.I2{613, 554}, vec.I2{667, 592}) },
		Dialogues: []awakengine.DialogueLine{
			{avatarAwakeman, "I don't think I can get past this fallen tree."},
		},
	},

	// See the W, hide it behind dialogue bubble >:)
	"was that a W": {
		Active: func() bool { return player.Pos().InRect(vec.I2{740, 390}, vec.I2{925, 444}) },
		Fire: func() {
			theW.pos = vec.I2{-100, -100}
			timeOfW = gameFrame
		},
		Dialogues: []awakengine.DialogueLine{
			{avatarAwakeman, `Wait, is that...!`},
			{avatarDucky, `QUACK! QUACK QUACK QUACK!`},
			{avatarAwakeman, `Quick!`},
		},
	},

	"it was?": {
		Depends: []string{"was that a W"},
		Active:  func() bool { return player.Pos().InRect(vec.I2{740, 400}, vec.I2{925, 444}) && gameFrame > timeOfW+120 },
		Dialogues: []awakengine.DialogueLine{
			{avatarAwakeman, `It was...it...uh...`},
			{avatarDucky, `QUaaaaaack.`},
			{avatarAwakeman, `I could have sworn something was hanging on that tree.`},
			{avatarNone, `*Ring*...`},
			{avatarAwakeman, `Wh--phone? Not when we're so close!`},
			{avatarNone, `*boop*`},
		},
	},

	"this way?": {
		Depends: []string{"it was?"},
		Active:  func() bool { return player.Pos().InRect(vec.I2{740, 440}, vec.I2{925, 472}) },
		Dialogues: []awakengine.DialogueLine{
			{avatarAwakeman, `Maybe it disappeared down this way.`},
		},
	},

	"holey trunk": {
		Active: func() bool { return player.Pos().InRect(vec.I2{849, 624}, vec.I2{912, 667}) },
		Dialogues: []awakengine.DialogueLine{
			{avatarDucky, `Quack!`},
			{avatarAwakeman, `I wonder what made this hole in this old tree trunk. Something round?`},
		},
	},

	"alamore": {
		// depends: []string{ "wasThatADoubleU" },
		Active: func() bool { return player.Pos().InRect(vec.I2{845, 853}, vec.I2{904, 898}) },
		Dialogues: []awakengine.DialogueLine{
			{avatarAwakeman, "Hey, Alamore!"},
			{avatarAlamore, "If it isn't Awakeman. Fancy seeing you out here! Enjoying the park?"},
			{avatarAwakeman, "It's been an interesting walk."},
			{avatarAlamore, "What brings you out here?"},
			{avatarAwakeman, `Well, I'm looking for a missing painting of a letter "W"--`},
			{avatarAlamore, "Surely a hero as brave and important as you has better things to do? :)"},
			{avatarAwakeman, `I'm not that important...`},
			{avatarAwakeman, `Um. I mean..."Awakeman" *is just* my name, I'm not a hero at all.`},
			{avatarAlamore, "No! Nonsense. You're smart, and very handsome, and a hero."},
			{avatarAlamore, "It's like you're an alien from another planet, you don't fit in here, not with the rest of the small-minded people who live in this town. :)"},
			{avatarAwakeman, `Um. Thanks.`},
			{avatarAlamore, "Who told you you're not a hero?"},
			{avatarAwakeman, "N--nobody?"},
			{avatarAlamore, "Oh come on, you can trust me. :)"},
			{avatarAwakeman, "No, I'm ... I don't think I'm special."},
			{avatarAlamore, "Hm. Well, we should have a chat about a very important project..."},
			{avatarNone, "AWAKEMAN! No. 27\nCREDITS\nProgramming, art, story, characters, concept\nJosh Deprez"},
			{avatarNone, "TUNE IN NEXT WEEK for the next GRIPPING EPISODE of AWAKEMAN!"},
			{avatarNone, "(It might be the end of this episode, but I won't stop you playing some more.)"},
		},
	},
}

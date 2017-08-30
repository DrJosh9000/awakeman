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
		FrameInfos: awakengine.BasicFrameInfos(7, -1, vec.I2{0, 0}),
		FrameSize:  vec.I2{34, 64},
	}

	avatarAwakeman        = &awakengine.SheetFrame{avatarsSheet, 0}
	avatarSJ              = &awakengine.SheetFrame{avatarsSheet, 1}
	avatarDucky           = &awakengine.SheetFrame{avatarsSheet, 2}
	avatarAlamore         = &awakengine.SheetFrame{avatarsSheet, 3}
	avatarAwakemanOnPhone = &awakengine.SheetFrame{avatarsSheet, 4}
	avatarBook            = &awakengine.SheetFrame{avatarsSheet, 5}
	avatarDoor            = &awakengine.SheetFrame{avatarsSheet, 6}

	exitTileCount = 0
	doorUnlocked  = false
)

func (g *Game) Triggers() []*awakengine.Trigger {
	if g.levelPreview || g.noTriggers {
		return nil
	}
	return []*awakengine.Trigger{
		// ------------------------ START ---------------------------
		{
			Name: "startGame",
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Text: "Awakeman! No. 40: Escape from the Dark Library\n\n(Click or tap to advance the dialogue.)", Slowness: -1},
					{Avatar: avatarAwakeman, Text: `*sigh*`, Slowness: 4},
					{Avatar: avatarDucky, Text: `Quack?`},
					{Avatar: avatarAwakeman, Text: `It's not you, Ducky.`},
					{Avatar: avatarAwakeman, Text: `I'm relieved that my ear isn't being talked off by Alamore anymore, but . . .`},
					{Avatar: avatarAwakeman, Text: `But now my phone is dead.`},
					{Avatar: avatarDucky, Text: `Quaaaaack.`},
					{Avatar: avatarAwakeman, Text: `Yeah. It was my light source.`},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 4},
					{Avatar: avatarAwakeman, Text: `Hey, don't suppose you have a built-in light of some kind?`},
					{Avatar: avatarDucky, Text: `Quack quack.`},
					{Text: "(Click or tap on things to move & interact.)"},
				}...)
			},
		},

		{
			Name:    "it's dark in here",
			Active:  func(f int) bool { return f > 20*12 },
			Depends: []string{"startGame"},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `My sight still hasn't adjusted to the dark, but I can make out things a few metres away.`},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 4},
					{Avatar: avatarAwakeman, Text: `Ducky, why are you here?`},
					{Avatar: avatarDucky, Text: `Quack!`},
					{Avatar: avatarAwakeman, Text: `Ask a silly question, get a silly answer.`},
					{Avatar: avatarDucky, Text: `*Ring, ring*`, Slowness: 2},
					{Avatar: avatarAwakeman, Text: `What?`},
					{Avatar: avatarSJ, Text: `Hello! Hello? Can you hear m--`, AutoNext: true},
					{Avatar: avatarAwakeman, Text: `Yeah, I can hear you. Science Jesus? You're talking to me through the duck?`},
					{Avatar: avatarSJ, Text: "Oh good, you picked up--Yes, I'm talking through the duck!\nAsk a silly question, get a silly answer."},
					{Avatar: avatarSJ, Text: "Where are you?"},
					{Avatar: avatarAwakeman, Text: "I'm in the office, it's dark and cold and my phone is dead and", AutoNext: true},
					{Avatar: avatarSJ, Text: "That's not . . . huh . . .", AutoNext: true, Slowness: 2},
					{Avatar: avatarAwakeman, Text: "and there's books everywhere and the power is out and", AutoNext: true},
					{Avatar: avatarSJ, Text: "I need to check on something. But I rang because I think you can find my painting nearby!"},
					{Avatar: avatarAwakeman, Text: ". . .", Slowness: 8},
					{Avatar: avatarAwakeman, Text: "Your *painting*?", Slowness: 2},
					{Avatar: avatarSJ, Text: `You're still going to find it for me, right? Sorry, gotta make like a tree and "split". Ciao!`},
					{Text: "*click*"},
					{Avatar: avatarDucky, Text: `Quack!`},
					{Avatar: avatarAwakeman, Text: `It's "make like a tree and l--" . . .`},
					{Avatar: avatarAwakeman, Text: `*sigh*`},
					{Avatar: avatarAwakeman, Text: `Who cares.`},
				}...)
			},
		},

		// ------------------------ STARTING AREA ---------------------------

		{
			Name:  "columns",
			Tiles: vec.RectRange(vec.I2{1, 12}, vec.I2{13, 13}),
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `These new columns are fancy.`},
					{Avatar: avatarAwakeman, Text: `. . .`},
					{Avatar: avatarAwakeman, Text: `Actually, they look pretty old.`},
					{Avatar: avatarAwakeman, Text: `Guess I never noticed them in the office before.`},
					{Avatar: avatarAwakeman, Text: `Something about having a ton of books nearby makes them stand out?`},
				}...)
			},
		},

		// ------------------------ BOOKS ---------------------------
		// ---------Top bookshelves near start ---------
		{
			Name:   "bookDiamondAge",
			Tiles:  []vec.I2{{2, 3}},
			Fire:   bookDiamondAge.offerToRead,
			Active: bookDiamondAge.available,
			Repeat: true,
		},
		{
			Name:   "bookDavidsonsPrinciples",
			Tiles:  []vec.I2{{4, 3}},
			Fire:   bookDavidsonsPrinciples.offerToRead,
			Active: bookDavidsonsPrinciples.available,
			Repeat: true,
		},
		{
			Name:  "bookFiftyShades",
			Tiles: []vec.I2{{5, 3}, {6, 3}, {7, 3}, {8, 3}, {9, 3}, {10, 3}, {11, 3}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `This seems to be the Fifty Shades of Grey section.`},
					{Avatar: avatarAwakeman, Text: `Like, a *lot* of copies of Fifty Shades . . .`},
					{Avatar: avatarAwakeman, Text: `I don't really feel like reading it right now, but there might be some more interesting-looking books elsewhere.`},
				}...)
			},
		},

		// --------- Lower left bookshelves ---------
		{
			Name:   "bookMakeGoodArt",
			Tiles:  []vec.I2{{2, 30}},
			Fire:   bookMakeGoodArt.offerToRead,
			Active: bookMakeGoodArt.available,
			Repeat: true,
		},
		{
			Name:   "bookGroupsOfAutomorphisms",
			Tiles:  []vec.I2{{4, 30}},
			Fire:   bookGroupsOfAutomorphisms.offerToRead,
			Active: bookGroupsOfAutomorphisms.available,
			Repeat: true,
		},
		{
			Name:   "book1984",
			Tiles:  []vec.I2{{6, 30}},
			Fire:   book1984.offerToRead,
			Active: book1984.available,
			Repeat: true,
		},
		{
			Name:   "bookHowToStaySane",
			Tiles:  []vec.I2{{8, 30}},
			Fire:   bookHowToStaySane.offerToRead,
			Active: bookHowToStaySane.available,
			Repeat: true,
		},
		{
			Name:   "bookLoseFriendsAndInfurate",
			Tiles:  []vec.I2{{10, 30}},
			Fire:   bookLoseFriendsAndInfurate.offerToRead,
			Active: bookLoseFriendsAndInfurate.available,
			Repeat: true,
		},

		// --------- Start of labyrinth ---------
		{
			Name:   "bookLibraryWizardLaw",
			Tiles:  []vec.I2{{20, 24}},
			Fire:   bookLibraryWizardLaw.offerToRead,
			Active: bookLibraryWizardLaw.available,
			Repeat: true,
		},
		{
			Name:   "bookContemporaryCryptology",
			Tiles:  []vec.I2{{22, 24}},
			Fire:   bookContemporaryCryptology.offerToRead,
			Active: bookContemporaryCryptology.available,
			Repeat: true,
		},

		// --------- Slightly into labyrinth ---------
		{
			Name:   "bookUnixInANutshell",
			Tiles:  []vec.I2{{19, 16}},
			Fire:   bookUnixInANutshell.offerToRead,
			Active: bookUnixInANutshell.available,
			Repeat: true,
		},
		{
			Name:   "bookSnuff",
			Tiles:  []vec.I2{{21, 16}},
			Fire:   bookSnuff.offerToRead,
			Active: bookSnuff.available,
			Repeat: true,
		},

		{
			Name:   "bookTrickyPeople",
			Tiles:  []vec.I2{{22, 9}},
			Fire:   bookTrickyPeople.offerToRead,
			Active: bookTrickyPeople.available,
			Repeat: true,
		},

		// --------- Just outside the spiral ---------
		{
			Name:   "bookHyperboleAndAHalf",
			Tiles:  []vec.I2{{27, 16}},
			Fire:   bookHyperboleAndAHalf.offerToRead,
			Active: bookHyperboleAndAHalf.available,
			Repeat: true,
		},
		{
			Name:   "bookRubaiyat",
			Tiles:  []vec.I2{{29, 16}},
			Fire:   bookRubaiyat.offerToRead,
			Active: bookRubaiyat.available,
			Repeat: true,
		},

		// --------- Out the spiral and up ---------
		{
			Name:   "bookProofsAndRefutations",
			Tiles:  []vec.I2{{28, 9}},
			Fire:   bookProofsAndRefutations.offerToRead,
			Active: bookProofsAndRefutations.available,
			Repeat: true,
		},
		{
			Name:   "bookiOSSwiftGameDevCookbook",
			Tiles:  []vec.I2{{30, 9}},
			Fire:   bookiOSSwiftGameDevCookbook.offerToRead,
			Active: bookiOSSwiftGameDevCookbook.available,
			Repeat: true,
		},

		// ------------------------ THE SPIRAL ---------------------------

		{
			Name:  "long corridor",
			Tiles: []vec.I2{{59, 16}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `A long corridor.`},
					{Avatar: avatarAwakeman, Text: `. . .`},
					{Avatar: avatarAwakeman, Text: `What a chore.`},
				}...)
			},
		},

		{
			Name:    "must proceed",
			Repeat:  true,
			Tiles:   []vec.I2{{59, 15}},
			Depends: []string{"long corridor"},
			Fire: func(int) {
				awakengine.PushDialogue(&awakengine.DialogueLine{Avatar: avatarAwakeman, Text: `There wasn't an exit back that way. I should press on . . .`})
				playerDelegate.SetPath([]vec.I2{{59*tileSize + 8, 17*tileSize + 8}})
			},
		},

		{
			Name:  "leading inwards",
			Tiles: []vec.I2{{38, 29}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `Oh . . . what?`},
					{Avatar: avatarAwakeman, Text: `More?`},
					{Avatar: avatarAwakeman, Text: `Just let it be over, let it be over`},
				}...)
			},
		},
		{
			Name:  "inwards 2",
			Tiles: []vec.I2{{39, 14}},
			Fire: func(int) {
				awakengine.PushDialogue(&awakengine.DialogueLine{
					Avatar: avatarAwakeman, Text: `Wha?`,
				})
			},
		},
		{
			Name:  "inwards 3",
			Tiles: []vec.I2{{53, 19}},
			Fire: func(int) {
				awakengine.PushDialogue(&awakengine.DialogueLine{
					Avatar: avatarAwakeman, Text: `Ha, ha.`,
				})
			},
		},
		{
			Name:  "before the centre",
			Tiles: []vec.I2{{45, 25}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `*blink*`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `*sob*`},
				}...)
			},
		},
		{
			Name:  "spiral centre",
			Tiles: []vec.I2{{47, 20}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `*sob*`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `Here I am, at the center of some kind of. . . a spiral.`},
					{Avatar: avatarAwakeman, Text: `I don't understand.`},
					{Avatar: avatarAwakeman, Text: `I don't understand any of it.`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `This is the most terrible thing it has ever been my misfortune to walk through.`},
					{Avatar: avatarAwakeman, Text: `Why would you make people walk almost the maximum possible distance?`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `What is the point?`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `I don't`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `I don't get it!`},
					{Avatar: avatarAwakeman, Text: `Not even government bureaucracy is this wasteful.`},
					{Avatar: avatarAwakeman, Text: `Why would you put this in the town council office?`},
					{Avatar: avatarAwakeman, Text: `. . .`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `Why?`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `Is there some secret here?`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `Am I going insane?`},
					{Avatar: avatarAwakeman, Text: `Am I imagining this?`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `*sob*`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `WHY?`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `*sob*`},
					{Avatar: avatarAwakeman, Text: `Is there no way out of here?`},
					{Avatar: avatarAwakeman, Text: `. . . . . . . .`, Slowness: 8},
					{Avatar: avatarDucky, Text: `*click*`},
					{Text: `(A recording starts to play.)`},
					{Avatar: avatarSJ, Text: `Awakeman. Ahem . . . `},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 4},
					{Avatar: avatarSJ, Text: `Listen.`},
					{Avatar: avatarSJ, Text: `I believe I've involved myself in something extremely dangerous.`},
					{Avatar: avatarSJ, Text: `And I've accidentally involved you, maybe.`},
					{Avatar: avatarSJ, Text: `If you're hearing this. . . oh, what a cliche.`},
					{Avatar: avatarSJ, Text: `Look, I'm probably dead.`},
					{Avatar: avatarSJ, Text: `Not likely to be a convenient time for you to hear this.`},
					{Avatar: avatarSJ, Text: `And I've accidentally involved you.`},
					{Avatar: avatarSJ, Text: `Well, you accidentally involved yourself . . . Unwittingly?`},
					{Avatar: avatarSJ, Text: `I have some leads.`},
					{Avatar: avatarSJ, Text: `I think you know what you need to find.`},
					{Avatar: avatarSJ, Text: `Sorry, if I say too much on this . . .`},
					{Avatar: avatarSJ, Text: `just. . .`},
					{Avatar: avatarSJ, Text: `Good luck,`},
					{Avatar: avatarSJ, Text: `my friend.`},
					{Avatar: avatarDucky, Text: `*click*`},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 2},
					{Avatar: avatarAwakeman, Text: `Great.`},
					{Avatar: avatarAwakeman, Text: `It's . . .`},
					{Avatar: avatarAwakeman, Text: `It's probably good that my feelings are totally numb right now.`},
					{Avatar: avatarAwakeman, Text: `. . . . . . . . . . . . . . . . . . . . . . . .`, Slowness: 4, AutoNext: true},
					{Avatar: avatarAwakeman, Text: `Shit.`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `Okay.`},
					{Avatar: avatarAwakeman, Text: `I don't have the luxury of staying in here anymore.`},
					{Avatar: avatarAwakeman, Text: `I don't have the luxury of wallowing in self-pity.`},
					{Avatar: avatarAwakeman, Text: ``},
					{Avatar: avatarAwakeman, Text: `Better get moving.`},
				}...)
			},
		},
		{
			Name:  "past the centre",
			Tiles: []vec.I2{{53, 22}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `What *is* this place?`},
					{Avatar: avatarAwakeman, Text: `If I figure out why everything seems out of place, I might be able to get my bearings.`},
					{Avatar: avatarAwakeman, Text: `. . . . . .`, Slowness: 8},
					{Avatar: avatarAwakeman, Text: `Wait . . .`, Slowness: 4},
					{Avatar: avatarAwakeman, Text: `This isn't the office, is it?`},
					{Avatar: avatarDucky, Text: `Quack!`},
					{Avatar: avatarAwakeman, Text: `. . . . . .`, Slowness: 8},
					{Avatar: avatarAwakeman, Text: `I'm in some kind of perverse library, with a godawful spiral corridor.`},
					{Avatar: avatarAwakeman, Text: `Books and building.`},
					{Avatar: avatarAwakeman, Text: `But not the town council office.`},
					{Avatar: avatarAwakeman, Text: `It *can't* be the office.`},
					{Avatar: avatarAwakeman, Text: `The past mayors were never this profligate.`},
				}...)
			},
		},

		{
			Name:  "past the centre 2",
			Tiles: []vec.I2{{40, 27}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: "The only sensible thing I can think of is . . ."},
					{Avatar: avatarAwakeman, Text: `I was teleported here, into this library, not the other way around.`},
					{Avatar: avatarDucky, Text: "Quack!"},
					{Avatar: avatarAwakeman, Text: `Unnngh. No, hang on . . .`},
					{Avatar: avatarAwakeman, Text: `I read a book back there, I think . . . Library Magic only physically affects books and shelves and such.`},
					{Avatar: avatarAwakeman, Text: `Not stick-figures.`},
					{Avatar: avatarAwakeman, Text: `Oh, and there was the fact that I got a call from Condiment, right as it happened, taking credit.`},
					{Avatar: avatarAwakeman, Text: `So the books must have been teleported on top of me!`},
					{Avatar: avatarAwakeman, Text: `But this isn't the office!`},
					{Avatar: avatarAwakeman, Text: `But I was in the office!`},
					{Avatar: avatarAwakeman, Text: `. . .`},
					{Avatar: avatarAwakeman, Text: `Argh!`},
					{Avatar: avatarAwakeman, Text: `What the hell is going on!`},
					{Avatar: avatarAwakeman, Text: `Having Science Jesus around would . . .`},
					{Avatar: avatarAwakeman, Text: ``},
				}...)
			},
		},

		{
			Name:  "spiral exit",
			Tiles: []vec.I2{{34, 17}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `The exit!`},
					{Avatar: avatarAwakeman, Text: `. . . from the spiral of complete pointlessness.`},
					{Avatar: avatarAwakeman, Text: `Why did I have to go through that?`},
				}...)
			},
		},

		{
			Name:    "not going back in there",
			Repeat:  true,
			Depends: []string{"spiral exit"},
			Tiles:   []vec.I2{{35, 17}},
			Fire: func(int) {
				awakengine.PushDialogue(&awakengine.DialogueLine{Avatar: avatarAwakeman, Text: `Not going back that way.`})
				playerDelegate.SetPath([]vec.I2{{33*tileSize + 8, 17*tileSize + 8}})
			},
		},

		// ------------------------ THE END ---------------------------

		{
			Name:  "about to hit the final area",
			Tiles: []vec.I2{{61, 14}},
		},

		{
			Name:    "don't leave the final area",
			Tiles:   []vec.I2{{61, 14}},
			Depends: []string{"about to hit the final area"},
			Repeat:  true,
			Fire: func(int) {
				awakengine.PushDialogue(&awakengine.DialogueLine{Avatar: avatarAwakeman, Text: `I don't feel like walking past the spiral yet again.`})
				playerDelegate.SetPath([]vec.I2{{61*tileSize + 8, 13*tileSize + 8}})
			},
		},

		{
			Name:   "exitDoorLocked",
			Tiles:  []vec.I2{{51, 3}},
			Repeat: true,
			Fire: func(int) {
				switch exitTileCount {
				case 0:
					awakengine.PushDialogue([]*awakengine.DialogueLine{
						{Avatar: avatarAwakeman, Text: `THE EXIT!`},
						{Avatar: avatarAwakeman, Text: `The real one!`},
						{Text: `(Awakeman rattles the door handle.)`},
						{Avatar: avatarAwakeman, Text: `It's locked!`},
					}...)
				case 1:
					awakengine.PushDialogue(&awakengine.DialogueLine{
						Avatar: avatarAwakeman, Text: `It's locked.`,
					})
				default:
					awakengine.PushDialogue(&awakengine.DialogueLine{
						Avatar: avatarAwakeman, Text: `It's still locked.`,
					})
				}
				exitTileCount++
			},
		},

		{
			Name:   "bookGLProject",
			Tiles:  []vec.I2{{56, 6}},
			Fire:   bookGLProject.offerToRead,
			Active: bookGLProject.available,
			Repeat: true,
		},

		{
			Name:    "booksWithSecrets",
			Depends: []string{"exitDoorLocked"},
			Tiles:   []vec.I2{{53, 6}, {54, 6}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					//{Avatar: avatarAwakeman, Text: `Endless maze of endless bookshelves.`},
					{Avatar: avatarAwakeman, Text: `Okay, let's try for a secret switch for the door.`},
					{Avatar: avatarAwakeman, Text: `If every fictional cliche is true, these things normally work via pulling the correct book out from the shelf.`},
					{Avatar: avatarAwakeman, Text: `Gonna visualise myself, out, walking around in the fresh air, and just . . . *feel* which is the right book.`},
					{Avatar: avatarAwakeman, Text: `Even though I'm not feeling much of anything anymore.`},
					{Avatar: avatarAwakeman, Text: `I choose you!`},
					{Text: `(Awakeman pulls on "The Secret" by Rhonda Byrne.)`},
					{Avatar: avatarBook, Text: `*nothing happens*`},
					{Avatar: avatarAwakeman, Text: `Darn. Okay.`},
					{Avatar: avatarAwakeman, Text: `Let's try again.`},
					{Text: `(Awakeman pulls on "The Secret of the Golden Flower" by Lu Dongbin as translated by Thomas Cleary.)`},
					{Avatar: avatarBook, Text: `*nothing happens*`},
					{Avatar: avatarAwakeman, Text: `Ahhhhh . . . this one.`},
					{Text: `(Awakeman pulls on "The Secret Life of Walter Mitty" by James Thurber.)`},
					{Avatar: avatarBook, Text: `*nothing happens*`},
					{Avatar: avatarAwakeman, Text: `Grargh!`},
					{Avatar: avatarAwakeman, Text: `Reveal your secrets!`},
					{Text: `(Awakeman pulls on "The Secret River" by Kate Grenville.)`},
					{Avatar: avatarBook, Text: `*nothing happens*`},
					{Avatar: avatarAwakeman, Text: `Argh!`},
					{Avatar: avatarAwakeman, Text: `Screw this bookshelf!`},
					{Text: `(In frustration, Awakeman kicks the shelf, and "Harry Potter and the Chamber of Secrets" by J. K. Rowling falls out . . . `},
					{Text: `(. . . revealing a small cavity in the shelf.)`},
					{Avatar: avatarAwakeman, Text: `*deep breath*`},
					{Avatar: avatarAwakeman, Text: `Sorry, Harry. I'm not a wizard.`},
					{Avatar: avatarAwakeman, Text: `But like you, I have arms.`},
					{Text: `(Awakeman reaches into the cavity and finds . . .)`},
					{Text: `(An old-fashioned key.)`},
					{Avatar: avatarAwakeman, Text: `*sob*!`},
					{Avatar: avatarAwakeman, Text: `Huzzah!`},
					{Avatar: avatarAwakeman, Text: `A completely useless key!`},
					{Text: `(Awakeman takes the key anyway.)`},
				}...)
			},
		},

		{
			Name:    "giveKey",
			Depends: []string{"booksWithSecrets"},
			Fire: func(int) {
				inventory.AddItems(&ItemKey{})
			},
		},

		{
			Name:   "rollCredits",
			Active: func(int) bool { return doorUnlocked },
			Fire: func(int) {
				scene.World.SetVisible(false)
				inventory.bubble.SetVisible(false)
				awakengine.PushDialogue(&awakengine.DialogueLine{
					Text: "Awakeman! No. 40 - Escape from the Dark Library\n\nBy Josh Deprez\n\nThanks for playing!",
				})
				gameOver = true
			},
		},
	}
}

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

	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

var (
	avatarsSheet = &awakengine.Sheet{
		Key:        "inv_avatars",
		FrameInfos: awakengine.BasicFrameInfos(6, -1, vec.I2{0, 0}),
		FrameSize:  vec.I2{34, 64},
	}

	avatarAwakeman        = &awakengine.SheetFrame{avatarsSheet, 0}
	avatarSJ              = &awakengine.SheetFrame{avatarsSheet, 1}
	avatarDucky           = &awakengine.SheetFrame{avatarsSheet, 2}
	avatarAlamore         = &awakengine.SheetFrame{avatarsSheet, 3}
	avatarAwakemanOnPhone = &awakengine.SheetFrame{avatarsSheet, 4}
	avatarBook            = &awakengine.SheetFrame{avatarsSheet, 5}

	bookEmpty = &book{
		title:  ``,
		author: ``,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `"`},
			{Avatar: avatarBook, Text: `""`},
		},
	}
	bookMakeGoodArt = &book{
		title:  `Neil Gaiman's "Make Good Art" Speech`,
		author: `Neil Gaiman`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `"If you're making mistakes, it means you're out there doing something.`},
			{Avatar: avatarBook, Text: `"And the mistakes in themselves can be useful. I once misspelled Caroline, in a letter, transposing the A and the O, and I thought, 'Coraline looks like a real name. . .'"`},
			{Avatar: avatarAwakeman, Text: `I don't know how relevant making good art is to doing good mayoring.`},
			{Avatar: avatarAwakeman, Text: `If I fail, then...what will happen to the town?`},
		},
	}
	bookTrickyPeople = &book{
		title:  `Tricky People`,
		author: `Andrew Fuller`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `"How to spot a Controller`},
			{Avatar: avatarBook, Text: `"This group contains the control freaks, the pedants, the outcome-obsessed meddlers, the end-justifies-the-means operators, the micro-managers and the nitpickers--and that's the most broad-minded of them.`},
			{Avatar: avatarBook, Text: `"Controllers are people who are happy if things are going their way.`},
			{Avatar: avatarBook, Text: `"Veer away from their wishes, however, and you will see an entirely different side of them."`},
			{Avatar: avatarAwakeman, Text: `. . .`},
			{Avatar: avatarAwakeman, Text: `Were Major Condiment and Library Wizard control freaks?`},
			{Avatar: avatarAwakeman, Text: `I took over and that's why I'm trapped in this library?`},
		},
	}
	bookHyperboleAndAHalf = &book{
		title:  `Hyperbole and a Half`,
		author: `Allie Brosh`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `"Depression Part One`},
			{Avatar: avatarBook, Text: `"Some people have a legitimate reason to feel depressed, but not me. I just woke up one day feeling arbitrarily sad and helpless.`},
			{Avatar: avatarBook, Text: `"It's disappointing to feel sad for no reason. Sadness can be alomst pleasantly indulgent when you have a way to justify it."`},
			{Avatar: avatarAwakeman, Text: `I'm pretty sure depression is a totally legitimate reason to feel depressed.`},
		},
	}
	bookHowToStaySane = &book{
		title:  `How to Stay Sane`,
		author: `Philippa Perry`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `"I have included this section on stories because a part of every successful therapy is about re-writing the narratives that define us, making new meaning and imagining different endings.`},
			{Avatar: avatarBook, Text: `"In the same way, part of staying sane is knowing what our story is and rewriting it when we need to."`},
		},
	}
	book1984 = &book{
		title:  `Nineteen Eighty-Four`,
		author: `George Orwell`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `"Again Winston's heart stirred painfully. It was inconceivable that this was anything other than a reference to Syme.`},
			{Avatar: avatarBook, Text: `"But Syme was not only dead, he was abolished, an unperson.`},
			{Avatar: avatarBook, Text: `"Any identifiable reference to him would have been mortally dangerous.`},
			{Avatar: avatarBook, Text: `"O'Brien's remark must obviously have been intended as a signal, a codeword.`},
			{Avatar: avatarBook, Text: `"By sharing a small act of thoughtcrime he had turned the two of them into accomplices."`},
			{Avatar: avatarAwakeman, Text: `Accomplices...`},
		},
	}
	bookLoseFriendsAndInfurate = &book{
		title:  `How to Lose Friends & Infuriate People`,
		author: `Jonar C. Nader`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `"New Wisdom #12`},
			{Avatar: avatarBook, Text: `"There's no smoke without fire, and no bad employee without a dubious manager."`},
			{Avatar: avatarAwakeman, Text: `This book is brightly-coloured, like fire.`},
			{Avatar: avatarDucky, Text: `Quack!`},
			{Avatar: avatarAwakeman, Text: `Packed full of wisdom?`},
		},
	}
	bookHollowFields = &book{
		title:  `Hollow Fields`,
		author: `Madeleine Rosca`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarAwakeman, Text: `It's a manga!`},
			{Avatar: avatarBook, Text: `"Hmmm...the coordinates just aren't right...Where could it be?`},
			{Avatar: avatarBook, Text: `"Doctor Bleak?`},
			{Avatar: avatarBook, Text: `"Were you...doing something? I heard a funny noise.`},
			{Avatar: avatarBook, Text: `"Hmm? I was just...asleep. Lack of sleep is starting to make you hear things, lass."`},
		},
	}
	bookGroupsOfAutomorphisms = &book{
		title:  `Groups of automorphisms of algebraic systems`,
		author: `B. I. Plotkin`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `"5.4.2. Left Engel groups. Right Engel elements.`},
			{Avatar: avatarBook, Text: `"A group in which all elements are left Engel elements is called a *left Engel group*.`},
			{Avatar: avatarBook, Text: `"Every locally nilpotent group is a left Engel group, but by a result of Golod the converse is not true.`},
			{Avatar: avatarBook, Text: `"The converse proposition can be proved under various additional assumptions."`},
			{Avatar: avatarAwakeman, Text: `I know what approximately zero of those sentences are saying.`},
		},
	}
	bookContemporaryCryptology = &book{
		title:  `Contemporary Cryptology`,
		author: `G. J. Simmons`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `"We have used the term 'subspace' imprecisely--and will continue to do so--to indicate a geometric object in an affine space that is either a subspace or is isomorphic to a subspace.`},
			{Avatar: avatarBook, Text: `"These objects are properly called *flats*, and only those flats that include the origin are actually subspaces."`},
			{Avatar: avatarAwakeman, Text: `Another theory-heavy book. It's also a physically heavy book.`},
		},
	}
	bookSnuff = &book{
		title:  `Snuff`,
		author: `Terry Pratchett`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `"Tears of the Mushroom picked up the plate and tentatively pushed it towards Vimes,`},
			{Avatar: avatarBook, Text: `"and said something that sounded like half a dozen coconuts rolling downstairs, but somehow managed to include the syllables *you* and *eat* and *I make*.`},
			{Avatar: avatarBook, Text: `"There seemed to be a pleading in her expression, as if trying to make him understand."`},
			{Avatar: avatarAwakeman, Text: `Reading in the dark is more fun than Mayoring!`},
		},
	}
	bookRubaiyat = &book{
		title:  `The Rubaiyat`,
		author: `Omar Khayyam, translated by Edward FitzGerald`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarAwakeman, Text: `It's a very small book.`},
			{Avatar: avatarBook, Text: `"XCIX. Ah, Love! could you and I with Him conspire`},
			{Avatar: avatarBook, Text: `"To grasp this sorry Scheme of Things entire,`},
			{Avatar: avatarBook, Text: `"Would not we shatter it to bits--and then`},
			{Avatar: avatarBook, Text: `"Re-mould it nearer to the Heart's Desire!"`},
		},
	}
	bookLibraryWizardLaw = &book{
		title:  `Advanced Library Magic, International Revised Ed.`,
		author: `Bernard K. Leftfiddle`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `"Fundamental Laws of Library Magic`},
			{Avatar: avatarAwakeman, Text: `Well well well. Maybe there's something in here that will help me get out of this library, or prevent having a whole spontaneous library crushing the office again!`},
			{Avatar: avatarBook, Text: "\"Preamble\nPrimary Tenets"},
			{Avatar: avatarBook, Text: `"1. A library wizard may not damage a book, or through inaction, cause a book to come to harm.`},
			{Avatar: avatarAwakeman, Text: `Duh. All they really care about is books, books, books!`},
			{Avatar: avatarBook, Text: `"2. The purpose of books and libraries is the dissemination of knowledge; the application of library magic more so.`},
			{Avatar: avatarBook, Text: `"3. Library magic cannot be used against patrons of the library, even for the purpose of extracting overdue fees or quieting rowdy readers."`},
			{Avatar: avatarAwakeman, Text: `That's . . . interesting.`},
			{Avatar: avatarAwakeman, Text: `If library magic can't be used against patrons, why am I so suddenly and magically buried under a ton of books?`},
		},
	}
)

type book struct {
	title    string
	author   string
	contents []*awakengine.DialogueLine
}

func (b *book) offerToRead(int) {
	awakengine.PushDialogue(&awakengine.DialogueLine{
		Text: "(Read a book?)",
		Buttons: []awakengine.ButtonSpec{
			{Label: "Read it", Action: b.read},
			{Label: "Leave it", Action: nil},
		},
	})
}

func (b *book) read() {
	awakengine.PushDialogue(b.contents...)
	awakengine.PushDialogueToBack(&awakengine.DialogueLine{
		Text: "(Take this book?)",
		Buttons: []awakengine.ButtonSpec{
			{Label: "Take it", Action: func() {
				awakengine.PushDialogue(
					&awakengine.DialogueLine{
						Text: fmt.Sprintf(`(Awakeman took "%s" by %s.)`, b.title, b.author),
					},
				)
			}},
			{Label: "Leave it", Action: func() {
				awakengine.PushDialogue(
					&awakengine.DialogueLine{
						Text: fmt.Sprintf(`(Awakeman put "%s" by %s back on the shelf.)`, b.title, b.author),
					},
				)
			}},
		},
	})
}

func (g *Game) Triggers() map[string]*awakengine.Trigger {
	if g.levelPreview || g.noTriggers {
		return nil
	}
	return map[string]*awakengine.Trigger{
		"startGame": {
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					/*
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
						}},*/
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
				}...)
			},
		},

		"it's dark in here": {
			Active:  func(f int) bool { return f > 20*6 },
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
				}...)
			},
		},

		// ------------------------ STARTING AREA ---------------------------

		// TODO: sequence about columns

		// ------------------------ BOOKS ---------------------------
		//
		// TODO: more books

		"bookMakeGoodArt": {
			Active: func(int) bool { return player.Pos.I2().InRect(vec.I2{X: 32, Y: 480}, vec.I2{X: 47, Y: 496}) },
			Fire:   bookMakeGoodArt.offerToRead,
		},
		"bookGroupsOfAutomorphisms": {
			Active: func(int) bool { return player.Pos.I2().InRect(vec.I2{X: 64, Y: 480}, vec.I2{X: 79, Y: 496}) },
			Fire:   bookGroupsOfAutomorphisms.offerToRead,
		},
		"book1984": {
			Active: func(int) bool { return player.Pos.I2().InRect(vec.I2{X: 96, Y: 480}, vec.I2{X: 111, Y: 496}) },
			Fire:   book1984.offerToRead,
		},

		"bookLibraryWizardLaw": {
			Active: func(int) bool { return player.Pos.I2().InRect(vec.I2{X: 304, Y: 384}, vec.I2{X: 320, Y: 400}) },
			Fire:   bookLibraryWizardLaw.offerToRead,
		},

		// ------------------------ THE SPIRAL ---------------------------

		"before the centre": {
			Active: func(int) bool { return player.Pos.I2().InRect(vec.I2{X: 675, Y: 396}, vec.I2{X: 691, Y: 405}) },
			Fire: func(int) {
				awakengine.PushDialogue(&awakengine.DialogueLine{
					Avatar: avatarAwakeman, Text: `How long is this going to go on for?`,
				})
			},
		},

		"spiral centre": {
			Active: func(int) bool { return player.Pos.I2().InRect(vec.I2{745, 325}, vec.I2{775, 351}) },
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
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
				}...)
			},
		},

		"past the centre": {
			Active: func(int) bool { return player.Pos.I2().InRect(vec.I2{712, 261}, vec.I2{725, 271}) },
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: "If I'm not in the office, where am I?"},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 8},
					{Avatar: avatarAwakeman, Text: "Some kind of perverse library. Books and building."},
				}...)
			},
		},

		"past the centre 2": {
			Active: func(int) bool { return player.Pos.I2().InRect(vec.I2{X: 915, Y: 315}, vec.I2{X: 924, Y: 323}) },
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: "Perhaps..."},
					{Avatar: avatarAwakeman, Text: "Maybe I was teleported here, not the other way around!"},
					{Avatar: avatarDucky, Text: "Quack!"},
				}...)
			},
		},

		// ------------------------ THE END ---------------------------
		//
	}
}

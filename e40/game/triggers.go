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
			{Avatar: avatarBook, Text: ``},
			{Avatar: avatarAwakeman, Text: ``},
		},
	}

	book1984 = &book{
		title:  `Nineteen Eighty-Four`,
		author: `George Orwell`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `Again Winston's heart stirred painfully. It was inconceivable that this was anything other than a reference to Syme.`},
			{Avatar: avatarBook, Text: `But Syme was not only dead, he was abolished, an unperson.`},
			{Avatar: avatarBook, Text: `Any identifiable reference to him would have been mortally dangerous.`},
			{Avatar: avatarBook, Text: `O'Brien's remark must obviously have been intended as a signal, a codeword.`},
			{Avatar: avatarBook, Text: `By sharing a small act of thoughtcrime he had turned the two of them into accomplices.`},
			{Avatar: avatarAwakeman, Text: `Accomplices...`},
		},
	}
	bookContemporaryCryptology = &book{
		title:  `Contemporary Cryptology`,
		author: `G. J. Simmons`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `We have used the term 'subspace' imprecisely--and will continue to do so--`},
			{Avatar: avatarBook, Text: `to indicate a geometric object in an affine space that is either a subspace or is isomorphic to a subspace.`},
			{Avatar: avatarBook, Text: `These objects are properly called *flats*, and only those flats that include the origin are actually subspaces.`},
			{Avatar: avatarAwakeman, Text: `Another theory-heavy book. It's also a physically heavy book.`},
		},
	}
	bookDavidsonsPrinciples = &book{
		title:  `Davidson's Principles and Practice of Medicine, 19th Ed.`,
		author: `Sir Stanley Davidson; ed. by Haslett, Chilvers, Boon, Colledge & Hunter`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarAwakeman, Text: `Oof! This is quite a hefty yellow tome.`},
			{Avatar: avatarBook, Text: "Acute Metabolic Complications\nHypoglycaemia"},
			{Avatar: avatarBook, Text: `Hypoglycaemia (blood glucose < 3.5 mmol/l) occurs often in diabetic patients treated with insulin but relatively infrequently in those taking a sulphonylurea drug.`},
			{Avatar: avatarBook, Text: `The risk of hypoglycaemia is the most important single factor limiting the attainment of the therapeutic goal, namely near-normal glycaemia.`},
			{Avatar: avatarAwakeman, Text: `Maybe I should get a check-up.`},
		},
	}
	bookDiamondAge = &book{
		title:  `The Diamond Age`,
		author: `Neal Stephenson`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `"Oh, I don't really know," Hackworth said.`},
			{Avatar: avatarBook, Text: `A movement caught his eye and he saw Fiona framed in a second-story window. She was undoing the latches, raising the sash.`},
			{Avatar: avatarBook, Text: `"I'm on a quest of sorts."`},
			{Avatar: avatarBook, Text: `"A quest for what, Mr. Hackworth?"`},
			{Avatar: avatarAwakeman, Text: `Welcome to my life.`},
		},
	}
	bookGroupsOfAutomorphisms = &book{
		title:  `Groups of automorphisms of algebraic systems`,
		author: `B. I. Plotkin`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `5.4.2. Left Engel groups. Right Engel elements.`},
			{Avatar: avatarBook, Text: `A group in which all elements are left Engel elements is called a *left Engel group*.`},
			{Avatar: avatarBook, Text: `Every locally nilpotent group is a left Engel group, but by a result of Golod the converse is not true.`},
			{Avatar: avatarBook, Text: `The converse proposition can be proved under various additional assumptions.`},
			{Avatar: avatarAwakeman, Text: `I know what approximately zero of those sentences are saying.`},
		},
	}
	bookHollowFields = &book{
		title:  `Hollow Fields`,
		author: `Madeleine Rosca`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarAwakeman, Text: `It's a manga!`},
			{Avatar: avatarAwakeman, Text: `Someone who runs this library has good taste!`},
			{Avatar: avatarBook, Text: `Hmmm...the coordinates just aren't right...Where could it be?`},
			{Avatar: avatarBook, Text: `Doctor Bleak?`},
			{Avatar: avatarBook, Text: `Were you...doing something? I heard a funny noise.`},
			{Avatar: avatarBook, Text: `Hmm? I was just...asleep. Lack of sleep is starting to make you hear things, lass.`},
		},
	}
	bookHowToStaySane = &book{
		title:  `How to Stay Sane`,
		author: `Philippa Perry`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `I have included this section on stories because a part of every successful therapy is about re-writing the narratives that define us,`},
			{Avatar: avatarBook, Text: `making new meaning and imagining different endings.`},
			{Avatar: avatarBook, Text: `In the same way, part of staying sane is knowing what our story is and rewriting it when we need to.`},
			{Avatar: avatarAwakeman, Text: `I can't exactly write myself out of being trapped in here.`},
		},
	}
	bookHyperboleAndAHalf = &book{
		title:  `Hyperbole and a Half`,
		author: `Allie Brosh`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `Depression Part One`},
			{Avatar: avatarBook, Text: `Some people have a legitimate reason to feel depressed, but not me. I just woke up one day feeling arbitrarily sad and helpless.`},
			{Avatar: avatarBook, Text: `It's disappointing to feel sad for no reason. Sadness can be alomst pleasantly indulgent when you have a way to justify it.`},
			{Avatar: avatarAwakeman, Text: `I'm pretty sure depression is a totally legitimate reason to feel depressed.`},
		},
	}
	bookiOSSwiftGameDevCookbook = &book{
		title:  `iOS Swift Game Development Cookbook`,
		author: `Jonathon Manning and Paris Buttfield-Addison`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: "Matrices\nA *matrix* is a grid of numbers (see Figure 8-1)."},
			{Avatar: avatarBook, Text: `On their own, matrices are just a way to store numbers. However, matrices are especially useful when they're combined with vectors.`},
			{Avatar: avatarBook, Text: `This is because you can multiply a matrix with a vector, which results in a changed version of the original vector.`},
			{Avatar: avatarAwakeman, Text: `Well, my phone is out of battery, so I won't be multiplying matrices and vectors on iOS anytime soon.`},
		},
	}
	bookLoseFriendsAndInfurate = &book{
		title:  `How to Lose Friends & Infuriate People`,
		author: `Jonar C. Nader`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `New Wisdom #12`},
			{Avatar: avatarBook, Text: `There's no smoke without fire, and no bad employee without a dubious manager.`},
			{Avatar: avatarAwakeman, Text: `This book is brightly-coloured, like fire.`},
			{Avatar: avatarDucky, Text: `Quack!`},
			{Avatar: avatarAwakeman, Text: `Packed full of wisdom?`},
			{Avatar: avatarAwakeman, Text: `I wonder if there's a book that will teach me how to speak "duck".`},
		},
	}
	bookMakeGoodArt = &book{
		title:  `Neil Gaiman's "Make Good Art" Speech`,
		author: `Neil Gaiman`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `If you're making mistakes, it means you're out there doing something.`},
			{Avatar: avatarBook, Text: `And the mistakes in themselves can be useful.`},
			{Avatar: avatarBook, Text: `I once misspelled Caroline, in a letter, transposing the A and the O, and I thought, "Coraline looks like a real name. . ."`},
			{Avatar: avatarAwakeman, Text: `I don't know how relevant making good art is to doing good mayoring, or good escaping-dark-libraries.`},
			{Avatar: avatarAwakeman, Text: `If I fail to be a good mayor, then what will happen to the town?`},
		},
	}
	bookProofsAndRefutations = &book{
		title:  `Proofs and Refutations`,
		author: `Imre Lakatos`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `Beta: I am glad that Sigma agreed with me, but I am afraid that I cannot agree with him.`},
			{Avatar: avatarBook, Text: `There are certainly three types of propositions: true ones, hopelessly false ones, and hopefully false ones.`},
			{Avatar: avatarBook, Text: `This last type can be improved into true propositions by adding a restrictive clause which states the exceptions.`},
			{Avatar: avatarAwakeman, Text: `I can't very well "prove" my way out of here.`},
		},
	}
	bookRubaiyat = &book{
		title:  `The Rubaiyat`,
		author: `Omar Khayyam, translated by Edward FitzGerald`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarAwakeman, Text: `It's a very small book.`},
			{Avatar: avatarBook, Text: `XCIX. Ah, Love! could you and I with Him conspire`},
			{Avatar: avatarBook, Text: `To grasp this sorry Scheme of Things entire,`},
			{Avatar: avatarBook, Text: `Would not we shatter it to bits--and then`},
			{Avatar: avatarBook, Text: `Re-mould it nearer to the Heart's Desire!`},
			{Avatar: avatarAwakeman, Text: `You got that right.`},
			{Avatar: avatarAwakeman, Text: `Stupid spiral.`},
		},
	}
	bookSnuff = &book{
		title:  `Snuff`,
		author: `Terry Pratchett`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `Tears of the Mushroom picked up the plate and tentatively pushed it towards Vimes,`},
			{Avatar: avatarBook, Text: `and said something that sounded like half a dozen coconuts rolling downstairs, but somehow managed to include the syllables *you* and *eat* and *I make*.`},
			{Avatar: avatarBook, Text: `There seemed to be a pleading in her expression, as if trying to make him understand.`},
			{Avatar: avatarAwakeman, Text: `Reading in the dark is more fun than Mayoring!`},
		},
	}
	bookTrickyPeople = &book{
		title:  `Tricky People`,
		author: `Andrew Fuller`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `How to spot a Controller`},
			{Avatar: avatarBook, Text: `This group contains the control freaks, the pedants, the outcome-obsessed meddlers, the end-justifies-the-means operators, the micro-managers and the nitpickers--`},
			{Avatar: avatarBook, Text: `and that's the most broad-minded of them.`},
			{Avatar: avatarBook, Text: `Controllers are people who are happy if things are going their way.`},
			{Avatar: avatarBook, Text: `Veer away from their wishes, however, and you will see an entirely different side of them.`},
			{Avatar: avatarAwakeman, Text: `. . .`},
			{Avatar: avatarAwakeman, Text: `Were Major Condiment and Library Wizard control freaks?`},
			{Avatar: avatarAwakeman, Text: `Seems most likely that this situation is retaliation over the bridge night cancellation . . .`},
			{Avatar: avatarAwakeman, Text: `. . . and that small matter of exiling them from town.`},
			{Avatar: avatarAwakeman, Text: `Well, it's working. I feel bad.`, Slowness: 4},
		},
	}
	bookUnixInANutshell = &book{
		title:  `Unix in a Nutshell, 4th Ed.`,
		author: `Arnold Robbins`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `shlock -f lockfile [-p PID] [-u] -[v]`},
			{Avatar: avatarBook, Text: `Create or verify a lockfile that can be used from shell scripts.`},
			{Avatar: avatarBook, Text: `shlock uses the rename(2) system call for making the final lock file; its operation is atomic.`},
			{Avatar: avatarAwakeman, Text: `Honestly, "shlock" seems more like a cartoon sound effect.`},
		},
	}

	bookLibraryWizardLaw = &book{
		title:  `Advanced Library Magic, International Revised Ed.`,
		author: `Bernard K. Leftfiddle`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `Chapter 1. Fundamental Laws of Library Magic`},
			{Avatar: avatarAwakeman, Text: `Aha! Maybe there's something in here that will help me get out of here, or prevent having a whole spontaneous library crushing the office again.`},
			{Avatar: avatarBook, Text: `Primary Tenets`},
			{Avatar: avatarBook, Text: `1. A library magic-user may not damage a book, or through inaction, cause a book to come to harm.`},
			{Avatar: avatarAwakeman, Text: `Makes sense. All they really care about is books anyway.`},
			{Avatar: avatarBook, Text: `2. The purpose of books and libraries is the dissemination of knowledge; the application of library magic more so.`},
			{Avatar: avatarBook, Text: `3. Library magic cannot be used against patrons of the library, even for the purpose of extracting overdue fees or quieting rowdy readers.`},
			{Avatar: avatarAwakeman, Text: `That's . . . `},
			{Avatar: avatarAwakeman, Text: `If library magic can't be used against patrons,`},
			{Avatar: avatarAwakeman, Text: `Why am I swimming in books?`},
			{Avatar: avatarAwakeman, Text: `Isn't this spontaneous library an example of weaponised library magic?`},
			{Avatar: avatarAwakeman, Text: `Maybe Library Wizard is in serious trouble with the library magic police!`},
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
						Text: fmt.Sprintf(`(Awakeman returned "%s" by %s to the shelf.)`, b.title, b.author),
					},
				)
			}},
		},
	})
}

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

		{
			Name:    "it's dark in here",
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
			Name:  "bookDiamondAge",
			Tiles: []vec.I2{{1, 3}, {2, 3}},
			Fire:  bookDiamondAge.offerToRead,
		},
		{
			Name:  "bookDavidsonsPrinciples",
			Tiles: []vec.I2{{3, 3}, {4, 3}},
			Fire:  bookDavidsonsPrinciples.offerToRead,
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
			Name:  "bookMakeGoodArt",
			Tiles: []vec.I2{{2, 30}, {3, 30}},
			Fire:  bookMakeGoodArt.offerToRead,
		},
		{
			Name:  "bookGroupsOfAutomorphisms",
			Tiles: []vec.I2{{4, 30}, {5, 30}},
			Fire:  bookGroupsOfAutomorphisms.offerToRead,
		},
		{
			Name:  "book1984",
			Tiles: []vec.I2{{6, 30}, {7, 30}},
			Fire:  book1984.offerToRead,
		},
		{
			Name:  "bookHowToStaySane",
			Tiles: []vec.I2{{8, 30}, {9, 30}},
			Fire:  bookHowToStaySane.offerToRead,
		},
		{
			Name:  "bookLoseFriendsAndInfurate",
			Tiles: []vec.I2{{10, 30}, {11, 30}},
			Fire:  bookLoseFriendsAndInfurate.offerToRead,
		},

		// --------- Start of labyrinth ---------
		{
			Name:  "bookLibraryWizardLaw",
			Tiles: []vec.I2{{19, 24}, {20, 24}},
			Fire:  bookLibraryWizardLaw.offerToRead,
		},
		{
			Name:  "bookContemporaryCryptology",
			Tiles: []vec.I2{{21, 24}, {22, 24}},
			Fire:  bookContemporaryCryptology.offerToRead,
		},

		// --------- Slightly into labyrinth ---------
		{
			Name:  "bookUnixInANutshell",
			Tiles: []vec.I2{{18, 16}, {19, 16}},
			Fire:  bookUnixInANutshell.offerToRead,
		},
		{
			Name:  "bookSnuff",
			Tiles: []vec.I2{{20, 16}, {21, 16}},
			Fire:  bookSnuff.offerToRead,
		},

		{
			Name:  "bookTrickyPeople",
			Tiles: []vec.I2{{21, 9}, {22, 9}},
			Fire:  bookTrickyPeople.offerToRead,
		},

		// --------- Just outside the spiral ---------
		{
			Name:  "bookHyperboleAndAHalf",
			Tiles: []vec.I2{{26, 16}, {27, 16}},
			Fire:  bookHyperboleAndAHalf.offerToRead,
		},
		{
			Name:  "bookRubaiyat",
			Tiles: []vec.I2{{28, 16}, {29, 16}},
			Fire:  bookRubaiyat.offerToRead,
		},

		// --------- Out the spiral and up ---------
		{
			Name:  "bookProofsAndRefutations",
			Tiles: []vec.I2{{27, 9}, {28, 9}},
			Fire:  bookProofsAndRefutations.offerToRead,
		},
		{
			Name:  "bookiOSSwiftGameDevCookbook",
			Tiles: []vec.I2{{29, 9}, {30, 9}},
			Fire:  bookiOSSwiftGameDevCookbook.offerToRead,
		},

		// ------------------------ THE SPIRAL ---------------------------

		{
			Name:  "long corridor",
			Tiles: []vec.I2{{59, 16}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `A long corridor. What's at the end?`},
					{Avatar: avatarAwakeman, Text: `I don't remember having such a long corridor in the office.`},
				}...)
			},
		},
		{
			Name:  "leading inwards",
			Tiles: []vec.I2{{38, 29}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `Oh . . . what?`},
					{Avatar: avatarAwakeman, Text: `More?`},
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
				awakengine.PushDialogue(&awakengine.DialogueLine{
					Avatar: avatarAwakeman, Text: `How long is this going to go on for?`,
				})
			},
		},
		{
			Name:  "spiral centre",
			Tiles: []vec.I2{{47, 20}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `What th`},
					{Avatar: avatarAwakeman, Text: `What is THE POINT of this spiralling corridor?`},
					{Avatar: avatarAwakeman, Text: `This is most terrible architecture it has ever been my misfortune to walk through!`},
					{Avatar: avatarAwakeman, Text: `Why would you make people walk almost the maximum possible distance?`},
					{Avatar: avatarAwakeman, Text: `What is the point?`},
					{Avatar: avatarAwakeman, Text: `I don't--I don't get it!`},
					{Avatar: avatarAwakeman, Text: `Not even government bureaucracy is this wasteful.`},
					{Avatar: avatarAwakeman, Text: `I . . .`},
					{Avatar: avatarAwakeman, Text: `I can't even`},
					{Avatar: avatarAwakeman, Text: `Why?`},
					{Avatar: avatarAwakeman, Text: `Is there some secret here?`},
					{Avatar: avatarAwakeman, Text: `Is it an altar used by some bizarre cult?`},
					{Avatar: avatarAwakeman, Text: `Why would you do put this in the town council office?`},
					{Avatar: avatarAwakeman, Text: `*sob*`},
					{Avatar: avatarAwakeman, Text: `WHY?`},
					{Avatar: avatarAwakeman, Text: `. . . . . . . .`, Slowness: 8},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 4},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 2},
					{Avatar: avatarAwakeman, Text: `This isn't the office,`, Slowness: 4},
					{Avatar: avatarAwakeman, Text: `. . . is it.`},
				}...)
			},
		},
		{
			Name:  "past the centre",
			Tiles: []vec.I2{{53, 22}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `If I'm not in the office, where am I?`},
					{Avatar: avatarAwakeman, Text: `. . .`, Slowness: 8},
					{Avatar: avatarAwakeman, Text: `Some kind of perverse library. Books and building. How did I get here?`},
				}...)
			},
		},

		{
			Name:  "past the centre 2",
			Tiles: []vec.I2{{40, 27}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: "The only sensible thing I can think of is . . ."},
					{Avatar: avatarAwakeman, Text: `I was teleported here, into this library, not the other way around!`},
					{Avatar: avatarDucky, Text: "Quack!"},
					{Avatar: avatarAwakeman, Text: `Who would do this?`},
					{Avatar: avatarAwakeman, Text: `Do you know the secrets of teleportation?`},
					{Avatar: avatarAwakeman, Text: `Argh!`},
					{Avatar: avatarAwakeman, Text: `Where's Science Jesus when you need them?`},
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

		// ------------------------ THE END ---------------------------
		{
			Name:  "booksWithSecrets",
			Tiles: []vec.I2{{53, 6}, {54, 6}},
			Fire: func(int) {
				awakengine.PushDialogue([]*awakengine.DialogueLine{
					{Avatar: avatarAwakeman, Text: `Endless maze of endless bookshelves.`},
					{Avatar: avatarAwakeman, Text: `Okay, let's try for a secret exit.`},
					{Avatar: avatarAwakeman, Text: `If every fictional cliche is true, these things normally work via pulling the correct book out from the shelf.`},
					{Avatar: avatarAwakeman, Text: `Gonna visualise myself, out, walking around in the fresh air, and just . . . *feel* which is the right book.`},
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
					{Avatar: avatarAwakeman, Text: `Screw this bookshelf!`},
					{Text: `(In frustration, Awakeman kicks the shelf, and "Harry Potter and the Chamber of Secrets" by J. K. Rowling falls out . . . `},
					{Text: `(. . . revealing a small cavity in the shelf.)`},
					{Avatar: avatarAwakeman, Text: `*deep breath*`},
					{Avatar: avatarAwakeman, Text: `Sorry, Harry. I'm not a wizard.`},
					{Avatar: avatarAwakeman, Text: `But I have arms.`},
					{Text: `(Awakeman reaches into the cavity and finds . . .)`},
					{Text: `(An old-fashioned key.)`},
					{Avatar: avatarAwakeman, Text: `*sob*!`},
					{Avatar: avatarAwakeman, Text: `Huzzah!`},
					{Avatar: avatarAwakeman, Text: `A completely useless key!`},
					{Text: `(Awakeman takes the key anyway.)`},
				}...)
			},
		},
	}
}

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
)

var (
	bookEmpty = &book{
		title:  ``,
		author: ``,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: ``},
			{Avatar: avatarAwakeman, Text: ``},
		},
	}

	// --------- REAL BOOKS ---------
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
		author: `Stanley Davidson, edited by Haslett, Chilvers, Boon, Colledge & Hunter`,
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
			{Avatar: avatarAwakeman, Text: `Someone has good taste.`},
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
			{Avatar: avatarBook, Text: `Some people have a legitimate reason to feel depressed, but not me.`},
			{Avatar: avatarBook, Text: `I just woke up one day feeling arbitrarily sad and helpless.`},
			{Avatar: avatarBook, Text: `It's disappointing to feel sad for no reason. Sadness can be almost pleasantly indulgent when you have a way to justify it.`},
			{Avatar: avatarAwakeman, Text: `I'm pretty sure depression is a totally legitimate reason to feel depressed.`},
			{Avatar: avatarAwakeman, Text: `Is being trapped in here a legitimate reason?`},
			{Avatar: avatarAwakeman, Text: `Is being told, by your probably-dead friend, to find a long-lost painting a legitimate reason?`},
			{Avatar: avatarAwakeman, Text: `Is having a dead phone and a companion that only quacks a legitimate reason?`},
			{Avatar: avatarAwakeman, Text: `Is being run down and smothered by your deputy a legitimate reason?`},
			{Avatar: avatarAwakeman, Text: `. . .`},
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
			{Avatar: avatarAwakeman, Text: `Re-mould the stupid spiral.`},
		},
	}
	bookSnuff = &book{
		title:  `Snuff`,
		author: `Terry Pratchett`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `Tears of the Mushroom picked up the plate and tentatively pushed it towards Vimes,`},
			{Avatar: avatarBook, Text: `and said something that sounded like half a dozen coconuts rolling downstairs, but somehow managed to include the syllables *you* and *eat* and *I make*.`},
			{Avatar: avatarBook, Text: `There seemed to be a pleading in her expression, as if trying to make him understand.`},
			{Avatar: avatarAwakeman, Text: `I should read in the dark more often!`},
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
			{Avatar: avatarAwakeman, Text: `Well, it's working . . .`},
			{Avatar: avatarAwakeman, Text: `I feel bad about doing that.`, Slowness: 4},
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

	// --------- FAKE BOOKS ---------
	bookLibraryWizardLaw = &book{
		title:  `Advanced Library Magic, International Revised Ed.`,
		author: `Bernard K. Leftfiddle`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarBook, Text: `Chapter 1. Fundamental Laws of Library Magic`},
			{Avatar: avatarAwakeman, Text: `Aha! Maybe there's something in here that will help me get out of here, or prevent having a whole spontaneous library crushing the office again.`},
			{Avatar: avatarBook, Text: `Primary Tenets`},
			{Avatar: avatarBook, Text: `1. A Library Magic-user may not damage a book, or through inaction, cause a book to come to harm.`},
			{Avatar: avatarAwakeman, Text: `Makes sense. All they really care about is books anyway.`},
			{Avatar: avatarBook, Text: `2. The purpose of books and libraries is the dissemination of knowledge; the application of Library Magic more so.`},
			{Avatar: avatarBook, Text: `3. In physical terms, Library Magic applies only to books, shelves, and other such items required for operating a library.`},
			{Avatar: avatarAwakeman, Text: `Which is why I'm blocked on all paths by books and not landmines.`},
			{Avatar: avatarBook, Text: `4. Library Magic cannot be used against patrons of the library, even for the purpose of extracting overdue fees or quieting rowdy readers.`},
			{Avatar: avatarAwakeman, Text: `That's . . . `},
			{Avatar: avatarAwakeman, Text: `If Library Magic can't be used against patrons, why am I swimming in books?`},
			{Avatar: avatarAwakeman, Text: `Isn't this spontaneous library an example of weaponised Library Magic?`},
			{Avatar: avatarAwakeman, Text: `Maybe Library Wizard is in serious trouble with the Library Magic police!`},
		},
	}

	bookGLProject = &book{
		title:  `G. L.`,
		author: `[REDACTED]`,
		contents: []*awakengine.DialogueLine{
			{Avatar: avatarAwakeman, Text: `It's the "G.L." folder I saw earlier.`},
			{Avatar: avatarAwakeman, Text: `But I walked all the way over here?`},
			{Avatar: avatarBook, Text: "G. L. PROJECT\n\nCONFIDENTIAL"},
			{Avatar: avatarBook, Text: `[REDACTED] [REDACTED] [REDACTED] [REDACTED] [REDACTED] [REDACTED] [REDACTED] [REDACTED]`},
			{Avatar: avatarAwakeman, Text: `It's all black rectangles and "redacteds".`},
			{Avatar: avatarAwakeman, Text: `Wait . . .`},
			{Avatar: avatarAwakeman, Text: `There's some kind of diagram?`},
			{Avatar: avatarAwakeman, Text: `I can't tell what it is.`},
			{Avatar: avatarAwakeman, Text: `Science Jesus probably c--. . .`},
			{Avatar: avatarAwakeman, Text: `*sob*`},
		},
	}
)

type book struct {
	title    string
	author   string
	contents []*awakengine.DialogueLine

	hasBeenRead bool
	offShelf    bool
}

func (b *book) available(int) bool {
	return !b.offShelf
}

func (b *book) offerToRead(int) {
	if b.hasBeenRead {
		b.offerToTake()
		return
	}
	awakengine.PushDialogue(&awakengine.DialogueLine{
		Text: "(Read a book?)",
		Buttons: []*awakengine.ButtonSpec{
			{Label: "Read it", Action: b.read},
			{Label: "Leave it", Action: nil},
		},
	})
}

func (b *book) read() {
	b.hasBeenRead = true
	b.offShelf = true
	awakengine.PushDialogue(b.contents...)
	b.offerToTake()
}

func (b *book) readWithoutCommentary() {
	for _, l := range b.contents {
		if l.Avatar != avatarBook {
			continue
		}
		awakengine.PushDialogueToBack(l)
	}
}

func (b *book) returnToShelf() {
	if !b.offShelf {
		return
	}
	b.offShelf = false
	awakengine.PushDialogueToBack(
		&awakengine.DialogueLine{
			Text: fmt.Sprintf(`(Awakeman returned "%s" by %s to the shelf.)`, b.title, b.author),
		},
	)
}

func (b *book) take() {
	// Can we fit the book?
	if inventory.NumItems() >= inventory.limit {
		awakengine.PushDialogue(&awakengine.DialogueLine{Avatar: avatarAwakeman, Text: `I can't carry any more stuff . . .`})
		b.returnToShelf()
		return
	}
	b.offShelf = true
	inventory.AddItems((*ItemBook)(b))
	awakengine.PushDialogue(&awakengine.DialogueLine{
		Text: fmt.Sprintf(`(Awakeman took "%s" by %s.)`, b.title, b.author),
	})
}

func (b *book) offerToTake() {
	takeQ := "(Take this book?)"
	if !b.offShelf {
		takeQ = fmt.Sprintf(`Take "%s"?`, b.title)
	}
	awakengine.PushDialogueToBack(&awakengine.DialogueLine{
		Text: takeQ,
		Buttons: []*awakengine.ButtonSpec{
			{Label: "Take it", Action: b.take},
			{Label: "Leave it", Action: b.returnToShelf},
		},
	})
}

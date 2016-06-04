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

var avatarsSheet = &awakengine.Sheet{
	Key:       "inv_avatars",
	Frames:    5,
	FrameSize: vec.I2{34, 64},
}

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
			Dialogues: []awakengine.DialogueLine{
				{Avatars: avatarsSheet, Index: avatarNone, Text: "Awakeman! No. 40: Escape from the Dark Library\n\n(Click or tap to advance the dialogue.)", Slowness: -1},
				{Avatars: avatarsSheet, Index: avatarAwakeman, Text: `*sigh*`, Slowness: 4},
				{Avatars: avatarsSheet, Index: avatarDucky, Text: `Quack?`},
				{Avatars: avatarsSheet, Index: avatarAwakeman, Text: `It's not you, Ducky. I'm glad my ear isn't being talked off by Alamore, but at the expense of my phone being dead.`},
				{Avatars: avatarsSheet, Index: avatarDucky, Text: `Quaaaaack.`},
				{Avatars: avatarsSheet, Index: avatarAwakeman, Text: `It was my light source.`},
				{Avatars: avatarsSheet, Index: avatarAwakeman, Text: `. . .`},
				{Avatars: avatarsSheet, Index: avatarAwakeman, Text: `Hey, don't suppose you have a light?`},
				{Avatars: avatarsSheet, Index: avatarDucky, Text: `Quack quack.`},
				{Avatars: avatarsSheet, Index: avatarNone, Text: "(Click or tap on things to move & interact.)"},
			},
		},

		"it's dark in here": {
			Depends: []string{"startGame"},
			Dialogues: []awakengine.DialogueLine{
				{Avatars: avatarsSheet, Index: avatarAwakeman, Text: `My sight seems to have adjusted to the dark.`},
				{Avatars: avatarsSheet, Index: avatarAwakeman, Text: `. . .`},
				{Avatars: avatarsSheet, Index: avatarAwakeman, Text: `Ducky, why are you here?`},
				{Avatars: avatarsSheet, Index: avatarDucky, Text: `Quack!`},
				{Avatars: avatarsSheet, Index: avatarAwakeman, Text: `Ask a silly question, get a silly answer.`},
				{Avatars: avatarsSheet, Index: avatarDucky, Text: `*bzzzt*`},
				{Avatars: avatarsSheet, Index: avatarAwakeman, Text: `What?`},
				{Avatars: avatarsSheet, Index: avatarSJ, Text: `Hello! Hello? Can you hear m--`, AutoNext: true},
				{Avatars: avatarsSheet, Index: avatarAwakeman, Text: `Yeah, I can hear you. Science Jesus? You're talking to me through the duck?`},
				{Avatars: avatarsSheet, Index: avatarSJ, Text: "Ask a silly question, get a silly answer.\nYes I'm talking through the duck!"},
			},
		},
	}
}

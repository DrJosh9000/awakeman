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
			Active: func(int) bool { return true },
			Dialogues: []awakengine.DialogueLine{
				//{avatarsAnim, avatarNone, "ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz 1234567890!@#$%%^&*();':\",./<>?"},
				{avatarsSheet, avatarNone, "Awakeman! No. 40: Escape from the Dark Library\n\n(Click or tap to advance the dialogue.)", nil},
				{avatarsSheet, avatarAwakeman, `*sigh*`, []string{"Demo", "Example"}},
				{avatarsSheet, avatarDucky, `Quack?`, nil},
				{avatarsSheet, avatarAwakeman, `It's not you, Ducky. I'm glad my ear isn't being talked off by Alamore, but at the expense of my phone being dead.`, nil},
				{avatarsSheet, avatarDucky, `Quaaaaack.`, nil},
				{avatarsSheet, avatarAwakeman, `It was my light source.`, nil},
				{avatarsSheet, avatarAwakeman, `. . .`, nil},
				{avatarsSheet, avatarAwakeman, `Hey, don't suppose you have a light?`, nil},
				{avatarsSheet, avatarDucky, `Quack quack.`, nil},
				{avatarsSheet, avatarNone, "(Click or tap on things to move & interact.)", nil},
			},
		},
	}
}

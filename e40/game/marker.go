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

var goalAckMarker = &ackMarker{
	anim: &awakengine.Anim{
		Key:       "inv_mark",
		Offset:    vec.I2{15, 15},
		Frames:    4,
		FrameSize: vec.I2{32, 32},
	},
}

type ackMarker struct {
	anim    *awakengine.Anim
	birth   int
	pos     vec.I2
	visible bool
}

func (a *ackMarker) begin(pos vec.I2, frame int) {
	a.birth = frame
	a.pos = pos
	a.visible = true
}

func (a *ackMarker) Anim() *awakengine.Anim { return a.anim }
func (a *ackMarker) Frame() int             { return 0 }
func (a *ackMarker) Pos() vec.I2            { return a.pos }

func (a *ackMarker) InWorld() bool { return true }
func (a *ackMarker) Retire() bool  { return false }
func (a *ackMarker) Visible() bool { return a.visible }
func (a *ackMarker) Z() int        { return a.pos.Y }
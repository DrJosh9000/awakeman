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

package common

import (
	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

type AckMarker struct {
	A *awakengine.Anim

	birth, frame int
	pos          vec.I2
	visible      bool
}

func (a *AckMarker) Begin(pos vec.I2, frame int) {
	a.birth = frame
	a.pos = pos
	a.visible = true
}

func (a *AckMarker) Anim() *awakengine.Anim { return a.A }
func (a *AckMarker) Frame() int             { return a.frame }
func (a *AckMarker) Pos() vec.I2            { return a.pos }

func (a *AckMarker) Update(t int) {
	if !a.visible {
		return
	}
	if t < a.birth || t >= a.birth+a.A.Frames {
		a.visible = false
		return
	}
	a.frame = t - a.birth
}

func (a *AckMarker) InWorld() bool { return true }
func (a *AckMarker) Retire() bool  { return false }
func (a *AckMarker) Visible() bool { return a.visible }
func (a *AckMarker) Z() int        { return a.pos.Y }

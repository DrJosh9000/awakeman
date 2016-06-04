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
	*awakengine.Sheet
	*awakengine.Playback
	awakengine.StaticOffset
	pos     vec.I2
	visible bool
}

func (a *AckMarker) Begin(pos vec.I2, frame int) {
	a.Playback.Reset()
	a.pos = pos
	a.visible = true
}

func (a *AckMarker) End()        { a.visible = false }
func (a *AckMarker) Pos() vec.I2 { return a.pos }

func (a *AckMarker) Update(int) {
	if !a.visible {
		return
	}
	a.Playback.Update(0)
}

func (a *AckMarker) Fixed() bool   { return false }
func (a *AckMarker) InWorld() bool { return true }
func (a *AckMarker) Retire() bool  { return false }
func (a *AckMarker) Visible() bool { return a.visible }
func (a *AckMarker) Z() int        { return -1 } // It's on the ground, which is -100

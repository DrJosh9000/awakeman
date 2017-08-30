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

import "github.com/DrJosh9000/awakengine"

type AckMarker struct {
	*awakengine.Sheet
}

func (a *AckMarker) Fixed(*awakengine.Sprite) bool                    { return false }
func (a *AckMarker) SpriteSheet(*awakengine.Sprite) *awakengine.Sheet { return a.Sheet }
func (a *AckMarker) Update(*awakengine.Sprite, int)                   {}
func (a *AckMarker) Z(*awakengine.Sprite) int                         { return -1 } // It's on the ground, which is -100, but below blocks, starting at 0.

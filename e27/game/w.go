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

var theW = &struct {
	*awakengine.StaticSprite
	*awakengine.StaticSemiobject
}{
	StaticSprite: &awakengine.StaticSprite{
		Sheet:          otherDoodadsSheet,
		StaticOffset:   awakengine.StaticOffset{16, 27},
		StaticPlayback: awakengine.StaticPlayback(1),
		P:              vec.I2{860, 453},
	},
	StaticSemiobject: &awakengine.StaticSemiobject{
		IW:  true,
		R:   false,
		V:   true,
		Zed: 453,
	},
}

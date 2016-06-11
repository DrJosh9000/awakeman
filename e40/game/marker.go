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
	"github.com/DrJosh9000/awakeman/common"
	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

var (
	goalAckMarker = &awakengine.Sprite{
		View: &awakengine.View{},
		SpriteDelegate: &common.AckMarker{
			Sheet: &awakengine.Sheet{
				Key:       "inv_mark",
				FrameSize: vec.I2{16, 16},
				FrameInfos: []awakengine.FrameInfo{
					// Start at 0, loop to 2.
					{Next: 1, Offset: vec.I2{8, 8}},
					{Next: 2, Offset: vec.I2{8, 8}},
					{Next: 3, Offset: vec.I2{8, 8}},
					{Next: 4, Offset: vec.I2{8, 8}},
					{Next: 5, Offset: vec.I2{8, 8}},
					{Next: 6, Offset: vec.I2{8, 8}},
					{Next: 2, Offset: vec.I2{8, 8}},
				},
			},
		},
	}
)

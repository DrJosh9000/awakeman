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

var player = &common.Player{
	P:  vec.F2{32 * 9, 32 * 10},
	UL: vec.I2{-3, -5},
	DR: vec.I2{4, 1},
	Anims: map[common.PlayerState]*awakengine.Anim{
		{common.PlayerActivityWalking, vec.Left}: {
			Key:       "walk_l",
			Offset:    vec.I2{8, 31},
			Frames:    7,
			FrameSize: vec.I2{16, 32},
		},
		{common.PlayerActivityWalking, vec.Right}: {
			Key:       "walk_r",
			Offset:    vec.I2{7, 31},
			Frames:    7,
			FrameSize: vec.I2{16, 32},
		},
		{common.PlayerActivityWalking, vec.Up}: {
			Key:       "walk_u",
			Offset:    vec.I2{7, 27},
			Frames:    14,
			FrameSize: vec.I2{16, 33},
		},
		{common.PlayerActivityWalking, vec.Down}: {
			Key:       "walk_d",
			Offset:    vec.I2{7, 27},
			Frames:    14,
			FrameSize: vec.I2{16, 33},
		},
		{common.PlayerActivityIdle, vec.Left}: {
			Key:       "idle_l",
			Offset:    vec.I2{7, 31},
			Frames:    1,
			FrameSize: vec.I2{16, 32},
		},
		{common.PlayerActivityIdle, vec.Up}: {
			Key:       "idle_l",
			Offset:    vec.I2{7, 31},
			Frames:    1,
			FrameSize: vec.I2{16, 32},
		},
		{common.PlayerActivityIdle, vec.Right}: {
			Key:       "idle_r",
			Offset:    vec.I2{9, 31},
			Frames:    1,
			FrameSize: vec.I2{16, 32},
		},
		{common.PlayerActivityIdle, vec.Down}: {
			Key:       "idle_r",
			Offset:    vec.I2{9, 31},
			Frames:    1,
			FrameSize: vec.I2{16, 32},
		},
	},
}

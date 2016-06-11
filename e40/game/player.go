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
	invIdleLeft = &awakengine.Sheet{
		Key: "inv_idle_l",
		FrameInfos: []awakengine.FrameInfo{
			{Next: 1, Duration: 100, Offset: vec.I2{7, 31}},
			{Next: 0, Duration: 20, Offset: vec.I2{7, 31}},
		},
		FrameSize: vec.I2{16, 32},
	}
	invIdleRight = &awakengine.Sheet{
		Key: "inv_idle_r",
		FrameInfos: []awakengine.FrameInfo{
			{Next: 1, Duration: 100, Offset: vec.I2{9, 31}},
			{Next: 0, Duration: 20, Offset: vec.I2{9, 31}},
		},
		FrameSize: vec.I2{16, 32},
	}

	playerAnims = map[common.PlayerState]*awakengine.Sheet{
		{common.PlayerActivityWalking, vec.Left}: {
			Key:        "inv_walk_l",
			FrameInfos: awakengine.BasicFrameInfos(7, 1, vec.I2{8, 31}),
			FrameSize:  vec.I2{16, 32},
		},
		{common.PlayerActivityWalking, vec.Right}: {
			Key:        "inv_walk_r",
			FrameInfos: awakengine.BasicFrameInfos(7, 1, vec.I2{7, 31}),
			FrameSize:  vec.I2{16, 32},
		},
		{common.PlayerActivityWalking, vec.Up}: {
			Key:        "inv_walk_u",
			FrameInfos: awakengine.BasicFrameInfos(14, 1, vec.I2{7, 27}),
			FrameSize:  vec.I2{16, 33},
		},
		{common.PlayerActivityWalking, vec.Down}: {
			Key:        "inv_walk_d",
			FrameInfos: awakengine.BasicFrameInfos(14, 1, vec.I2{7, 27}),
			FrameSize:  vec.I2{16, 33},
		},
		{common.PlayerActivityIdle, vec.Left}:  invIdleLeft,
		{common.PlayerActivityIdle, vec.Up}:    invIdleLeft,
		{common.PlayerActivityIdle, vec.Right}: invIdleRight,
		{common.PlayerActivityIdle, vec.Down}:  invIdleRight,
	}

	playerDelegate = &common.Player{
		UL:    vec.I2{-3, -5},
		DR:    vec.I2{4, 1},
		Anims: playerAnims,
		//State:  common.PlayerState{common.PlayerActivityIdle, vec.Right},
	}

	player = &awakengine.Sprite{
		Pos:            vec.F2{16*8 + 8, 16*4 + 8},
		SpriteDelegate: playerDelegate,
	}
)

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

const playerSpeed = 1

var (
	playerAnims = map[PlayerState]*awakengine.Anim{
		{playerWalking, vec.Left}: {
			Key:       "walk_l",
			Offset:    vec.I2{8, 31},
			Frames:    7,
			FrameSize: vec.I2{16, 32},
			Mode:      awakengine.AnimLoop,
		},
		{playerWalking, vec.Right}: {
			Key:       "walk_r",
			Offset:    vec.I2{7, 31},
			Frames:    7,
			FrameSize: vec.I2{16, 32},
			Mode:      awakengine.AnimLoop,
		},
		{playerWalking, vec.Up}: {
			Key:       "walk_u",
			Offset:    vec.I2{7, 27},
			Frames:    14,
			FrameSize: vec.I2{16, 33},
			Mode:      awakengine.AnimLoop,
		},
		{playerWalking, vec.Down}: {
			Key:       "walk_d",
			Offset:    vec.I2{7, 27},
			Frames:    14,
			FrameSize: vec.I2{16, 33},
			Mode:      awakengine.AnimLoop,
		},
		{playerIdle, vec.Left}: {
			Key:       "idle_l",
			Offset:    vec.I2{7, 31},
			Frames:    1,
			FrameSize: vec.I2{16, 32},
			Mode:      awakengine.AnimLoop,
		},
		{playerIdle, vec.Up}: {
			Key:       "idle_l",
			Offset:    vec.I2{7, 31},
			Frames:    1,
			FrameSize: vec.I2{16, 32},
			Mode:      awakengine.AnimLoop,
		},
		{playerIdle, vec.Right}: {
			Key:       "idle_r",
			Offset:    vec.I2{9, 31},
			Frames:    1,
			FrameSize: vec.I2{16, 32},
			Mode:      awakengine.AnimLoop,
		},
		{playerIdle, vec.Down}: {
			Key:       "idle_r",
			Offset:    vec.I2{9, 31},
			Frames:    1,
			FrameSize: vec.I2{16, 32},
			Mode:      awakengine.AnimLoop,
		},
	}

	player = &Player{
		pos: vec.F2{32 * 9, 32 * 10},
	}
	playerUL = vec.I2{-3, -5}
	playerDR = vec.I2{4, 1}
)

// PlayerState describes the current simple state of the player: what are they doing,
// and what direction are they facing.
type PlayerState struct {
	a PlayerActivity
	d vec.Direction
}

// PlayerActivity describes what the player is doing.
type PlayerActivity int

const (
	playerIdle = PlayerActivity(iota)
	playerWalking
)

// Player encapsulates all the state of the player character.
type Player struct {
	pos   vec.F2
	frame int
	path  []vec.I2
	state PlayerState
}

// Anim implements awakengine.Sprite.
func (p *Player) Anim() *awakengine.Anim { return playerAnims[p.state] }

// Frame implements awakengine.Sprite.
func (p *Player) Frame() int { return p.frame }

// Pos implements awakengine.Sprite.
func (p *Player) Pos() vec.I2 { return p.pos.I2() }

// Footprint implements awakengine.Unit.
func (p *Player) Footprint() (ul, dr vec.I2) { return playerUL, playerDR }

// GoIdle implements awakengine.Unit.
func (p *Player) GoIdle() {
	p.frame = 0
	p.state.a = playerIdle
	p.path = nil
}

// Path implements awakengine.Unit.
func (p *Player) Path() []vec.I2 { return p.path }

// Update implements awakengine.Unit.
func (p *Player) Update(frame int, event awakengine.Event) {
	if event.Type == awakengine.EventMouseUp {
		c := event.Pos
		goalAckMarker.Birth = frame
		goalAckMarker.P = c
		p.path = awakengine.Navigate(p.Pos(), c)
	}

	if len(p.path) > 0 {
		d := p.path[0].F2().Sub(p.pos)
		np := d.Norm()

		// If at the current point, either aim at the next point, or stop.
		if np < 1 {
			if len(p.path) > 1 {
				p.pos = p.path[0].F2() // Important to stop the player clipping through the wall...
				p.path = p.path[1:]
				d = p.path[0].F2().Sub(p.pos)
				np = d.Norm()
			} else {
				p.pos = p.path[0].F2()
				p.GoIdle()
				return
			}
		}
		p.state.a = playerWalking
		p.state.d = d.Dir()
		if frame%animPeriod == 0 {
			// Player walks straight there.
			p.pos = p.pos.Add(d.Mul(playerSpeed / np))
			p.frame++
		}
	} else {
		p.GoIdle()
	}
}

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

const playerSpeed = 1

// PlayerState describes the current simple state of the player: what are they doing,
// and what direction are they facing.
type PlayerState struct {
	A PlayerActivity
	D vec.Direction
}

// PlayerActivity describes what the player is doing.
type PlayerActivity int

const (
	PlayerActivityIdle = PlayerActivity(iota)
	PlayerActivityWalking
)

type Anim struct {
	*awakengine.Sheet
	awakengine.StaticOffset
	*awakengine.Playback
}

// Player encapsulates all the state of the player character.
type Player struct {
	// Templatey stuff
	Anims  map[PlayerState]*Anim
	UL, DR vec.I2

	// Instancey stuff
	*Anim
	P     vec.F2
	path  []vec.I2
	state PlayerState
}

func (p *Player) SetPath(path []vec.I2) { p.path = path }

func (p *Player) Fixed() bool   { return false }
func (p *Player) InWorld() bool { return true }
func (p *Player) Retire() bool  { return false }
func (p *Player) Visible() bool { return true }
func (p *Player) Z() int        { return p.Pos().Y }

func (p *Player) Pos() vec.I2                { return p.P.I2() }
func (p *Player) Footprint() (ul, dr vec.I2) { return p.UL, p.DR }

func (p *Player) GoIdle() {
	p.state.A = PlayerActivityIdle
	p.Anim = p.Anims[p.state]
	p.Anim.Playback.Reset()
	p.path = nil
}

func (p *Player) Path() []vec.I2 { return p.path }

func (p *Player) Update(t int) {
	if len(p.path) == 0 {
		if p.state.A != PlayerActivityIdle {
			p.GoIdle()
		}
		p.Anim.Update(t)
		return
	}

	// Take the next pixel step
	d := p.path[0].F2().Sub(p.P)
	np := d.Norm()

	// If at the current point, either aim at the next point, or stop.
	if np < 1 {
		if len(p.path) > 1 {
			p.P = p.path[0].F2() // Important to stop the player clipping through the wall...
			p.path = p.path[1:]
			d = p.path[0].F2().Sub(p.P)
			np = d.Norm()
		} else {
			p.P = p.path[0].F2()
			p.GoIdle()
			return
		}
	}
	p.state.A = PlayerActivityWalking
	p.state.D = d.Dir()
	p.Anim = p.Anims[p.state]

	// Player walks straight there.
	p.P = p.P.Add(d.Mul(playerSpeed / np))
	p.Anim.Update(t)
}

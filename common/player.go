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

// PlayerActivity describes what the player is doing.
type PlayerActivity int

const (
	PlayerActivityIdle = PlayerActivity(iota)
	PlayerActivityWalking
)

// PlayerState describes the current simple state of the player: what are they doing,
// and what direction are they facing.
type PlayerState struct {
	A PlayerActivity
	D vec.Direction
}

// Player encapsulates all the state of the player character.
type Player struct {
	// Templatey stuff
	Anims     map[PlayerState]*awakengine.Sheet
	UL, DR    vec.I2
	AckMarker *awakengine.Sprite

	// Instancey stuff
	path  []vec.I2
	state PlayerState
}

func (p *Player) SetPath(path []vec.I2) {
	p.path = path
	if len(path) == 0 {
		p.GoIdle()
		return
	}
	p.AckMarker.ResetAnim()
	p.AckMarker.Pos = path[len(path)-1].F2()
	p.AckMarker.SetVisible(true)
}

func (p *Player) Footprint() (ul, dr vec.I2) { return p.UL, p.DR }

func (p *Player) GoIdle() {
	p.state.A = PlayerActivityIdle
	//p.Sprite.ResetAnim()
	p.path = nil
	p.AckMarker.SetVisible(false)
}

func (p *Player) Path() []vec.I2 { return p.path }

func (p *Player) Fixed(*awakengine.Sprite) bool                    { return false }
func (p *Player) SpriteSheet(*awakengine.Sprite) *awakengine.Sheet { return p.Anims[p.state] }
func (p *Player) Z(s *awakengine.Sprite) int                       { return s.Pos.I2().Y }

func (p *Player) Update(s *awakengine.Sprite, _ int) {
	if len(p.path) == 0 {
		if p.state.A != PlayerActivityIdle {
			p.GoIdle()
			s.ResetAnim()
		}
		return
	}

	// Take the next pixel step
	d := p.path[0].F2().Sub(s.Pos)
	np := d.Norm()

	// If at the current point, either aim at the next point, or stop.
	if np < 1 {
		if len(p.path) > 1 {
			s.Pos = p.path[0].F2() // Important to stop the player clipping through the wall...
			p.path = p.path[1:]
			d = p.path[0].F2().Sub(s.Pos)
			np = d.Norm()
		} else {
			s.Pos = p.path[0].F2()
			p.GoIdle()
			s.ResetAnim()
			return
		}
	}
	p.state.A = PlayerActivityWalking
	p.state.D = d.Dir()

	// Player walks straight there.
	s.Pos = s.Pos.Add(d.Mul(playerSpeed / np))
}

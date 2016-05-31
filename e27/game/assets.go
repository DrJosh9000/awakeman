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
	common "github.com/DrJosh9000/awakeman/common/assets"
	"github.com/DrJosh9000/awakeman/e27/assets"
	"github.com/DrJosh9000/awakengine"
)

func init() {
	// Black-on-white assets.
	awakengine.RegisterImage("avatars", common.AvatarsPNG)
	awakengine.RegisterImage("bubble", common.BubblePNG)
	awakengine.RegisterImage("idle_l", common.PlayerIdleLPNG)
	awakengine.RegisterImage("idle_r", common.PlayerIdleRPNG)
	awakengine.RegisterImage("mark", common.MarkPNG)
	awakengine.RegisterImage("munro", common.MunroPNG)
	awakengine.RegisterImage("walk_d", common.PlayerWalkDPNG)
	awakengine.RegisterImage("walk_l", common.PlayerWalkLPNG)
	awakengine.RegisterImage("walk_r", common.PlayerWalkRPNG)
	awakengine.RegisterImage("walk_u", common.PlayerWalkUPNG)

	// e27 assets.
	awakengine.RegisterImage("doodads", assets.DoodadsPNG)
	awakengine.RegisterImage("level1", assets.Level1PNG)
	awakengine.RegisterImage("tiles", assets.TilesPNG)
	awakengine.RegisterImage("trees", assets.TreesPNG)
}

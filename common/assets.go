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
	"github.com/DrJosh9000/awakeman/common/assets"
	"github.com/DrJosh9000/awakengine"
)

func RegisterAssets() {
	// Black-on-white assets.
	awakengine.RegisterImage("avatars", assets.AvatarsPNG)
	awakengine.RegisterImage("bubble", assets.BubblePNG)
	awakengine.RegisterImage("idle_l", assets.PlayerIdleLPNG)
	awakengine.RegisterImage("idle_r", assets.PlayerIdleRPNG)
	awakengine.RegisterImage("mark", assets.MarkPNG)
	awakengine.RegisterImage("munro", assets.MunroPNG)
	awakengine.RegisterImage("walk_d", assets.PlayerWalkDPNG)
	awakengine.RegisterImage("walk_l", assets.PlayerWalkLPNG)
	awakengine.RegisterImage("walk_r", assets.PlayerWalkRPNG)

	// White-on-black assets.
	awakengine.RegisterImage("inv_avatars", assets.InvAvatarsPNG)
	awakengine.RegisterImage("inv_bubble", assets.InvBubblePNG)
	awakengine.RegisterImage("inv_idle_l", assets.InvPlayerIdleLPNG)
	awakengine.RegisterImage("inv_idle_r", assets.InvPlayerIdleRPNG)
	awakengine.RegisterImage("inv_mark", assets.InvMarkPNG)
	awakengine.RegisterImage("inv_munro", assets.InvMunroPNG)
	awakengine.RegisterImage("inv_walk_d", assets.InvPlayerWalkDPNG)
	awakengine.RegisterImage("inv_walk_l", assets.InvPlayerWalkLPNG)
	awakengine.RegisterImage("inv_walk_r", assets.InvPlayerWalkRPNG)
	awakengine.RegisterImage("inv_walk_u", assets.InvPlayerWalkUPNG)
}

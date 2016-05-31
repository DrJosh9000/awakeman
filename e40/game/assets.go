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
	"github.com/DrJosh9000/awakeman/e40/assets"
	"github.com/DrJosh9000/awakengine"
)

func init() {
	// White-on-black assets.
	awakengine.RegisterImage("inv_avatars", common.InvAvatarsPNG)
	awakengine.RegisterImage("inv_bubble", common.InvBubblePNG)
	awakengine.RegisterImage("inv_idle_l", common.InvPlayerIdleLPNG)
	awakengine.RegisterImage("inv_idle_r", common.InvPlayerIdleRPNG)
	awakengine.RegisterImage("inv_mark", common.InvMarkPNG)
	awakengine.RegisterImage("inv_munro", common.InvMunroPNG)
	awakengine.RegisterImage("inv_walk_d", common.InvPlayerWalkDPNG)
	awakengine.RegisterImage("inv_walk_l", common.InvPlayerWalkLPNG)
	awakengine.RegisterImage("inv_walk_r", common.InvPlayerWalkRPNG)
	awakengine.RegisterImage("inv_walk_u", common.InvPlayerWalkUPNG)

	// Black-on-white assets.
	awakengine.RegisterImage("bubble", common.BubblePNG)
	awakengine.RegisterImage("munro", common.MunroPNG)

	// e40 assets.
	awakengine.RegisterImage("level2blockmap", assets.Level2BlockMapPNG)
	awakengine.RegisterImage("libraryblocks", assets.LibraryBlocksPNG)
}

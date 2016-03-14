package game

import (
	"github.com/DrJosh9000/awakeman/e27/assets"
	"github.com/DrJosh9000/awakengine"
)

func init() {
	awakengine.RegisterImage("avatars", assets.AvatarsPNG)
	awakengine.RegisterImage("bubble", assets.BubblePNG)
	awakengine.RegisterImage("doodads", assets.DoodadsPNG)
	awakengine.RegisterImage("idle_l", assets.PlayerIdleLPNG)
	awakengine.RegisterImage("idle_r", assets.PlayerIdleRPNG)
	awakengine.RegisterImage("level1", assets.Level1PNG)
	awakengine.RegisterImage("mark", assets.MarkPNG)
	awakengine.RegisterImage("munro", assets.MunroPNG)
	awakengine.RegisterImage("tiles", assets.TilesPNG)
	awakengine.RegisterImage("trees", assets.TreesPNG)
	awakengine.RegisterImage("walk_d", assets.PlayerWalkDPNG)
	awakengine.RegisterImage("walk_l", assets.PlayerWalkLPNG)
	awakengine.RegisterImage("walk_r", assets.PlayerWalkRPNG)
	awakengine.RegisterImage("walk_u", assets.PlayerWalkUPNG)
}

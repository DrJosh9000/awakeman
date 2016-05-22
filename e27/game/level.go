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

const (
	tileSize    = 32 // pixels on the edge of a square tile
	blockHeight = 0  // Introduced in a later engine.

	level1Key  = "level1"
	tileMapKey = "tiles"
)

type level struct{}

// Doodads provides objects above the base, that can be flattened onto the terrain
// most of the time.
func (level) Doodads() []*awakengine.Doodad {
	return []*awakengine.Doodad{
		// Above the north building.
		{BaseDoodad: baseDoodads["tree1"], P: vec.I2{122, 200}},
		{BaseDoodad: baseDoodads["tree2"], P: vec.I2{150, 80}},
		{BaseDoodad: baseDoodads["tree1"], P: vec.I2{220, 85}},

		// Next to the west lake.
		{BaseDoodad: baseDoodads["tree2"], P: vec.I2{76, 319}},
		{BaseDoodad: baseDoodads["tree1"], P: vec.I2{53, 357}},

		// In the clearing
		{BaseDoodad: baseDoodads["tree3"], P: vec.I2{557, 498}},
		{BaseDoodad: baseDoodads["puffyplant"], P: vec.I2{437, 467}},
		{BaseDoodad: baseDoodads["puffyplant"], P: vec.I2{470, 561}},
		{BaseDoodad: baseDoodads["puffyplant"], P: vec.I2{432, 581}},
		{BaseDoodad: baseDoodads["puffyplant"], P: vec.I2{554, 591}},
		{BaseDoodad: baseDoodads["puffyplant"], P: vec.I2{653, 470}},
		{BaseDoodad: baseDoodads["puffyplant"], P: vec.I2{612, 436}},

		// Between the buildings.
		{BaseDoodad: baseDoodads["tree1"], P: vec.I2{259, 654}},
		{BaseDoodad: baseDoodads["tree1"], P: vec.I2{319, 706}},
		{BaseDoodad: baseDoodads["tree2"], P: vec.I2{261, 740}},

		// Blockage from clearing.
		{BaseDoodad: baseDoodads["tree9"], P: vec.I2{627, 574}},

		// Along the path along the north.
		{BaseDoodad: baseDoodads["bench_h"], P: vec.I2{500, 80}},
		{BaseDoodad: baseDoodads["bench_h"], P: vec.I2{580, 80}},
		{BaseDoodad: baseDoodads["tree5"], P: vec.I2{471, 141}},
		{BaseDoodad: baseDoodads["tree5"], P: vec.I2{511, 245}},
		{BaseDoodad: baseDoodads["tree5"], P: vec.I2{576, 145}},
		{BaseDoodad: baseDoodads["tree4"], P: vec.I2{610, 250}},
		{BaseDoodad: baseDoodads["tree5"], P: vec.I2{692, 148}},
		{BaseDoodad: baseDoodads["tree7"], P: vec.I2{760, 243}},

		// Toward the south-west.
		{BaseDoodad: baseDoodads["tree1"], P: vec.I2{78, 643}},
		{BaseDoodad: baseDoodads["tree2"], P: vec.I2{99, 704}},
		{BaseDoodad: baseDoodads["tree1"], P: vec.I2{158, 773}},
		{BaseDoodad: baseDoodads["tree2"], P: vec.I2{121, 939}},
		{BaseDoodad: baseDoodads["tree4"], P: vec.I2{183, 944}},

		// Trees in the east.
		{BaseDoodad: baseDoodads["tree1"], P: vec.I2{932, 285}},
		{BaseDoodad: baseDoodads["tree2"], P: vec.I2{934, 365}},
		{BaseDoodad: baseDoodads["tree8"], P: vec.I2{885, 645}},
		{BaseDoodad: baseDoodads["tree2"], P: vec.I2{937, 720}},
		{BaseDoodad: baseDoodads["tree7"], P: vec.I2{765, 958}},
		{BaseDoodad: baseDoodads["tree2"], P: vec.I2{669, 838}},
		{BaseDoodad: baseDoodads["tree4"], P: vec.I2{732, 831}},

		// The W tree.
		//theW,
		{BaseDoodad: baseDoodads["tree1"], P: vec.I2{860, 452}},

		// End-game in south-east.
		{BaseDoodad: baseDoodads["alamore"], P: vec.I2{877, 877}},
	}
}

// Source is the paletted PNG to use as the base terrain layer - pixel at (x,y) becomes
// the tile at (x,y).
func (level) Source() string {
	return level1Key
}

// TileInfos maps indexes to information about the terrain.
func (level) TileInfos() []awakengine.TileInfo {
	return []awakengine.TileInfo{
		{Name: "grass1", Block: false}, // 0
		{Name: "grass2", Block: false},
		{Name: "square2", Block: true},
		{Name: "cliff", Block: true},
		{Name: "black", Block: true},
		{Name: "grass_edge_d", Block: false}, // 5
		{Name: "black_edge_u", Block: true},
		{Name: "brick_wall", Block: true},
		{Name: "brick_wall_edge_l", Block: true},
		{Name: "brick_wall_edge_r", Block: true},
		{Name: "water_edge_r", Block: true}, // 10
		{Name: "water_edge_l", Block: true},
		{Name: "grey", Block: true},
		{Name: "grass_corner_rd", Block: false},
		{Name: "grass_corner_ld", Block: false},
		{Name: "hole", Block: true}, //15
		{Name: "pyramid", Block: true},
		{Name: "wire_fence", Block: true},
		{Name: "water_edge_ulcc", Block: true},
		{Name: "water_edge_ulcv", Block: true},
		{Name: "water_edge_u", Block: true}, //20
		{Name: "water_edge_urcv", Block: true},
		{Name: "water", Block: true},
		{Name: "water_edge_urcc", Block: true},
		{Name: "tall_grass", Block: false},
		{Name: "path_corner_dl", Block: false}, //25
		{Name: "path_corner_dr", Block: false},
		{Name: "path_corner_ur", Block: false},
		{Name: "path_corner_ul", Block: false},
		{Name: "path_h", Block: false},
		{Name: "path_v", Block: false}, //30
		{Name: "path_d", Block: false},
		{Name: "path_r", Block: false},
		{Name: "path_l", Block: false},
		{Name: "path_u", Block: false},
		{Name: "roof_dl", Block: true}, //35
		{Name: "roof_ul", Block: true},
		{Name: "roof_dr", Block: true},
		{Name: "roof_ur", Block: true},
		{Name: "roof_l", Block: true},
		{Name: "roof_r", Block: true}, //40
		{Name: "roof_d", Block: true},
		{Name: "roof_u", Block: true},
		{Name: "roof", Block: true},
		{Name: "water_edge_drcc", Block: true},
		{Name: "water_edge_d", Block: true}, //45
		{Name: "water_edge_dlcc", Block: true},
		{Name: "trees1", Block: true},
		{Name: "trees2", Block: true},
		{Name: "trees1_u", Block: true},
		{Name: "trees2_u", Block: true}, //50
		{Name: "trees1_d", Block: true},
		{Name: "trees2_d", Block: true},
		{Name: "pathend_r", Block: false},
		{Name: "pathend_d", Block: false},
	}
}

// Tiles is an image containing square tiles, and the size of each tile.
func (level) Tiles() (string, int, int) {
	return tileMapKey, tileSize, blockHeight
}

func (*Game) Level() awakengine.Level {
	return level{}
}

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

func (*Game) Level() (*awakengine.Level, error) {
	tm, sz, err := awakengine.ImageAsMap("level1")
	if err != nil {
		return nil, err
	}
	return &awakengine.Level{
		Doodads: []*awakengine.Doodad{
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
			{BaseDoodad: baseDoodads["tree1"], P: vec.I2{860, 452}},

			// End-game in south-east.
			{BaseDoodad: baseDoodads["alamore"], P: vec.I2{877, 877}},
		},
		MapSize: sz,
		TileMap: tm,
		TileInfos: []awakengine.TileInfo{
			{Name: "black", Blocking: true}, // 0
			{Name: "grass1", Blocking: false},
			{Name: "grass2", Blocking: false},
			{Name: "square2", Blocking: true},
			{Name: "cliff", Blocking: true},
			{Name: "grass_edge_d", Blocking: false}, // 5
			{Name: "black_edge_u", Blocking: true},
			{Name: "brick_wall", Blocking: true},
			{Name: "brick_wall_edge_l", Blocking: true},
			{Name: "brick_wall_edge_r", Blocking: true},
			{Name: "water_edge_r", Blocking: true}, // 10
			{Name: "water_edge_l", Blocking: true},
			{Name: "grey", Blocking: true},
			{Name: "grass_corner_rd", Blocking: false},
			{Name: "grass_corner_ld", Blocking: false},
			{Name: "hole", Blocking: true}, //15
			{Name: "pyramid", Blocking: true},
			{Name: "wire_fence", Blocking: true},
			{Name: "water_edge_ulcc", Blocking: true},
			{Name: "water_edge_ulcv", Blocking: true},
			{Name: "water_edge_u", Blocking: true}, //20
			{Name: "water_edge_urcv", Blocking: true},
			{Name: "water", Blocking: true},
			{Name: "water_edge_urcc", Blocking: true},
			{Name: "tall_grass", Blocking: false},
			{Name: "path_corner_dl", Blocking: false}, //25
			{Name: "path_corner_dr", Blocking: false},
			{Name: "path_corner_ur", Blocking: false},
			{Name: "path_corner_ul", Blocking: false},
			{Name: "path_h", Blocking: false},
			{Name: "path_v", Blocking: false}, //30
			{Name: "path_d", Blocking: false},
			{Name: "path_r", Blocking: false},
			{Name: "path_l", Blocking: false},
			{Name: "path_u", Blocking: false},
			{Name: "roof_dl", Blocking: true}, //35
			{Name: "roof_ul", Blocking: true},
			{Name: "roof_dr", Blocking: true},
			{Name: "roof_ur", Blocking: true},
			{Name: "roof_l", Blocking: true},
			{Name: "roof_r", Blocking: true}, //40
			{Name: "roof_d", Blocking: true},
			{Name: "roof_u", Blocking: true},
			{Name: "roof", Blocking: true},
			{Name: "water_edge_drcc", Blocking: true},
			{Name: "water_edge_d", Blocking: true}, //45
			{Name: "water_edge_dlcc", Blocking: true},
			{Name: "trees1", Blocking: true},
			{Name: "trees2", Blocking: true},
			{Name: "trees1_u", Blocking: true},
			{Name: "trees2_u", Blocking: true}, //50
			{Name: "trees1_d", Blocking: true},
			{Name: "trees2_d", Blocking: true},
			{Name: "pathend_r", Blocking: false},
			{Name: "pathend_d", Blocking: false},
		},
		TilesetKey:  "tiles",
		TileSize:    32,
		BlockHeight: 0,
		Obstacles:   precomputedObstacles,
		Paths:       precomputedPaths,
	}, nil
}

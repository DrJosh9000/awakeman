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
	"fmt"

	"github.com/DrJosh9000/awakengine"
	//	"github.com/DrJosh9000/vec"
)

func (*Game) Level() (*awakengine.Level, error) {
	tm, sz, err := awakengine.ImageAsMap("level2tilemap")
	if err != nil {
		return nil, err
	}
	bm, sz2, err := awakengine.ImageAsMap("level2blockmap")
	if err != nil {
		return nil, err
	}
	if sz != sz2 {
		return nil, fmt.Errorf("map sizes not equal [%v != %v]", sz, sz2)
	}

	return &awakengine.Level{
		Doodads:  nil,
		MapSize:  sz,
		TileMap:  tm,
		BlockMap: bm,
		TileInfos: []awakengine.TileInfo{
			{Name: "black"},
		},
		BlockInfos: []awakengine.TileInfo{
			{Name: "black"},
			{Name: "prism", Blocking: true},
		},
		TilesetKey:  "black16",
		BlocksetKey: "prism",
		TileSize:    16,
		BlockHeight: 32,
	}, nil
}

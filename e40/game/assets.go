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
	"github.com/DrJosh9000/awakeman/common"
	"github.com/DrJosh9000/awakeman/e40/assets"
	"github.com/DrJosh9000/awakengine"
)

func init() {
	common.RegisterWoBAssets()
	awakengine.RegisterImage("black16", assets.Black16PNG)
	awakengine.RegisterImage("level2tilemap", assets.Level2TileMapPNG)
	awakengine.RegisterImage("level2blockmap", assets.Level2BlockMapPNG)
	awakengine.RegisterImage("librarytiles", assets.LibraryTilesPNG)
	awakengine.RegisterImage("prism", assets.PrismPNG)
}

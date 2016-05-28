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

package main

import (
	"flag"
	"log"

	"github.com/DrJosh9000/awakeman/e40/game"
	"github.com/DrJosh9000/awakengine"
)

var (
	cfg          = &awakengine.Config{}
	levelPreview = flag.Bool("levelpreview", false, "Draw a huge level, but have no triggers")
)

func init() {
	flag.BoolVar(&cfg.Debug, "debug", false, "Enables debug visualisations")
	flag.StringVar(&cfg.LevelGeomDump, "levelgeom", "", "If set, dumps computed obstacles & paths graphs to this file")
	flag.StringVar(&cfg.RecordingFile, "r", "", "If set, records a GIF of the screen to this file")
	flag.IntVar(&cfg.RecordingFrames, "frames", 120, "The number of frames to record into the GIF")
}

func main() {
	flag.Parse()
	g := game.New(*levelPreview)
	if err := awakengine.Run(g, cfg); err != nil {
		log.Fatalf("Cannot run game: %v", err)
	}
}

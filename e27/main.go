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

	"github.com/DrJosh9000/awakeman/e27/game"
	"github.com/DrJosh9000/awakengine"
)

var (
	debug           = flag.Bool("debug", false, "Enables debug visualisations.")
	levelPreview    = flag.Bool("levelpreview", false, "Draw a huge level, but have no triggers.")
	recordingFile   = flag.String("r", "", "If set, records a GIF of the screen to this file.")
	recordingFrames = flag.Int("frames", 120, "The number of frames to record into the GIF.")
)

func main() {
	flag.Parse()
	if err := awakengine.Run(game.New(*levelPreview), *debug, *recordingFile, *recordingFrames); err != nil {
		log.Fatalf("Cannot run game: %v", err)
	}
}

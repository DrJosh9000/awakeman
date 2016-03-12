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
	awakengine.Init(game.New(*levelPreview), *debug)
	if err := awakengine.Load(); err != nil {
		log.Fatalf("Cannot load game: %v", err)
	}
	if err := awakengine.Run(*recordingFile, *recordingFrames); err != nil {
		log.Fatalf("Cannot run game: %v", err)
	}
}

package game

import (
	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

var baseDoodads = map[string]*awakengine.BaseDoodad{
	// --------------------- TREES ---------------------
	"tree1": {
		Anim: &awakengine.Anim{
			Key:       "trees",
			Offset:    vec.I2{30, 59},
			Frames:    9,
			FrameSize: vec.I2{64, 64},
		},
		Frame: 0,
		UL:    vec.I2{25, 56},
		DR:    vec.I2{34, 63},
	},
	"tree2": {
		Anim: &awakengine.Anim{
			Key:       "trees",
			Offset:    vec.I2{33, 61},
			Frames:    9,
			FrameSize: vec.I2{64, 64},
		},
		Frame: 1,
		UL:    vec.I2{29, 55},
		DR:    vec.I2{38, 63},
	},
	"tree3": {
		Anim: &awakengine.Anim{
			Key:       "trees",
			Offset:    vec.I2{9, 58},
			Frames:    9,
			FrameSize: vec.I2{64, 64},
		},
		Frame: 2,
		UL:    vec.I2{3, 53},
		DR:    vec.I2{60, 62},
	},
	"tree4": {
		Anim: &awakengine.Anim{
			Key:       "trees",
			Offset:    vec.I2{31, 60},
			Frames:    9,
			FrameSize: vec.I2{64, 64},
		},
		Frame: 3,
		UL:    vec.I2{4, 53},
		DR:    vec.I2{59, 61},
	},
	"tree5": {
		Anim: &awakengine.Anim{
			Key:       "trees",
			Offset:    vec.I2{31, 59},
			Frames:    9,
			FrameSize: vec.I2{64, 64},
		},
		Frame: 4,
		UL:    vec.I2{19, 52},
		DR:    vec.I2{44, 63},
	},
	"tree6": {
		Anim: &awakengine.Anim{
			Key:       "trees",
			Offset:    vec.I2{31, 56},
			Frames:    9,
			FrameSize: vec.I2{64, 64},
		},
		Frame: 5,
		UL:    vec.I2{22, 51},
		DR:    vec.I2{34, 57},
	},
	"tree7": {
		Anim: &awakengine.Anim{
			Key:       "trees",
			Offset:    vec.I2{32, 58},
			Frames:    9,
			FrameSize: vec.I2{64, 64},
		},
		Frame: 6,
		UL:    vec.I2{19, 52},
		DR:    vec.I2{44, 63},
	},
	"tree8": {
		Anim: &awakengine.Anim{
			Key:       "trees",
			Offset:    vec.I2{31, 59},
			Frames:    9,
			FrameSize: vec.I2{64, 64},
		},
		Frame: 7,
		UL:    vec.I2{18, 56},
		DR:    vec.I2{41, 63},
	},
	"tree9": {
		Anim: &awakengine.Anim{
			Key:       "trees",
			Offset:    vec.I2{7, 59},
			Frames:    9,
			FrameSize: vec.I2{64, 64},
		},
		Frame: 8,
		UL:    vec.I2{3, 53},
		DR:    vec.I2{63, 61},
	},

	// --------------------- SMALL DOODADS ---------------------
	"bench_h": {
		Anim: &awakengine.Anim{
			Key:       "doodads",
			Offset:    vec.I2{16, 24},
			Frames:    6,
			FrameSize: vec.I2{32, 32},
		},
		Frame: 0,
		UL:    vec.I2{1, 15},
		DR:    vec.I2{30, 24},
	},
	"W": {
		Anim: &awakengine.Anim{
			Key:       "doodads",
			Offset:    vec.I2{16, 27},
			Frames:    6,
			FrameSize: vec.I2{32, 32},
		},
		Frame: 1,
		UL:    vec.I2{3, 23},
		DR:    vec.I2{26, 29},
	},
	"puffyplant": {
		Anim: &awakengine.Anim{
			Key:       "doodads",
			Offset:    vec.I2{16, 24},
			Frames:    6,
			FrameSize: vec.I2{32, 32},
		},
		Frame: 2,
		UL:    vec.I2{6, 24},
		DR:    vec.I2{24, 27},
	},
	"evilplant": {
		Anim: &awakengine.Anim{
			Key:       "doodads",
			Offset:    vec.I2{16, 24},
			Frames:    6,
			FrameSize: vec.I2{32, 32},
		},
		Frame: 3,
		UL:    vec.I2{2, 22},
		DR:    vec.I2{30, 27},
	},
	"bench_v": {
		Anim: &awakengine.Anim{
			Key:       "doodads",
			Offset:    vec.I2{8, 22},
			Frames:    6,
			FrameSize: vec.I2{32, 32},
		},
		Frame: 4,
		UL:    vec.I2{4, 13},
		DR:    vec.I2{15, 30},
	},
	"alamore": {
		Anim: &awakengine.Anim{
			Key:       "doodads",
			Offset:    vec.I2{16, 25},
			Frames:    6,
			FrameSize: vec.I2{32, 32},
		},
		Frame: 5,
		UL:    vec.I2{8, 20},
		DR:    vec.I2{23, 28},
	},
}

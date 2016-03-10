package game

var (
	tileInfo = []awakengine.TileInfo{
		{name: "grass1", block: false}, // 0
		{name: "grass2", block: false},
		{name: "square2", block: true},
		{name: "cliff", block: true},
		{name: "black", block: true},
		{name: "grass_edge_d", block: false}, // 5
		{name: "black_edge_u", block: true},
		{name: "brick_wall", block: true},
		{name: "brick_wall_edge_l", block: true},
		{name: "brick_wall_edge_r", block: true},
		{name: "water_edge_r", block: true}, // 10
		{name: "water_edge_l", block: true},
		{name: "grey", block: true},
		{name: "grass_corner_rd", block: false},
		{name: "grass_corner_ld", block: false},
		{name: "hole", block: true}, //15
		{name: "pyramid", block: true},
		{name: "wire_fence", block: true},
		{name: "water_edge_ulcc", block: true},
		{name: "water_edge_ulcv", block: true},
		{name: "water_edge_u", block: true}, //20
		{name: "water_edge_urcv", block: true},
		{name: "water", block: true},
		{name: "water_edge_urcc", block: true},
		{name: "tall_grass", block: false},
		{name: "path_corner_dl", block: false}, //25
		{name: "path_corner_dr", block: false},
		{name: "path_corner_ur", block: false},
		{name: "path_corner_ul", block: false},
		{name: "path_h", block: false},
		{name: "path_v", block: false}, //30
		{name: "path_d", block: false},
		{name: "path_r", block: false},
		{name: "path_l", block: false},
		{name: "path_u", block: false},
		{name: "roof_dl", block: true}, //35
		{name: "roof_ul", block: true},
		{name: "roof_dr", block: true},
		{name: "roof_ur", block: true},
		{name: "roof_l", block: true},
		{name: "roof_r", block: true}, //40
		{name: "roof_d", block: true},
		{name: "roof_u", block: true},
		{name: "roof", block: true},
		{name: "water_edge_drcc", block: true},
		{name: "water_edge_d", block: true}, //45
		{name: "water_edge_dlcc", block: true},
		{name: "trees1", block: true},
		{name: "trees2", block: true},
		{name: "trees1_u", block: true},
		{name: "trees2_u", block: true}, //50
		{name: "trees1_d", block: true},
		{name: "trees2_d", block: true},
		{name: "pathend_r", block: false},
		{name: "pathend_d", block: false},
	}

	theW = &Doodad{
		BaseDoodad: baseDoodads["W"],
		pos:        vec.I2{860, 453},
	}

	doodads = []*awakengine.Doodad{
		// Above the north building.
		{BaseDoodad: baseDoodads["tree1"], pos: vec.I2{122, 200}},
		{BaseDoodad: baseDoodads["tree2"], pos: vec.I2{150, 80}},
		{BaseDoodad: baseDoodads["tree1"], pos: vec.I2{220, 85}},

		// Next to the west lake.
		{BaseDoodad: baseDoodads["tree2"], pos: vec.I2{76, 319}},
		{BaseDoodad: baseDoodads["tree1"], pos: vec.I2{53, 357}},

		// In the clearing
		{BaseDoodad: baseDoodads["tree3"], pos: vec.I2{557, 498}},
		{BaseDoodad: baseDoodads["puffyplant"], pos: vec.I2{437, 467}},
		{BaseDoodad: baseDoodads["puffyplant"], pos: vec.I2{470, 561}},
		{BaseDoodad: baseDoodads["puffyplant"], pos: vec.I2{432, 581}},
		{BaseDoodad: baseDoodads["puffyplant"], pos: vec.I2{554, 591}},
		{BaseDoodad: baseDoodads["puffyplant"], pos: vec.I2{653, 470}},
		{BaseDoodad: baseDoodads["puffyplant"], pos: vec.I2{612, 436}},

		// Between the buildings.
		{BaseDoodad: baseDoodads["tree1"], pos: vec.I2{259, 654}},
		{BaseDoodad: baseDoodads["tree1"], pos: vec.I2{319, 706}},
		{BaseDoodad: baseDoodads["tree2"], pos: vec.I2{261, 740}},

		// Blockage from clearing.
		{BaseDoodad: baseDoodads["tree9"], pos: vec.I2{627, 574}},

		// Along the path along the north.
		{BaseDoodad: baseDoodads["bench_h"], pos: vec.I2{500, 80}},
		{BaseDoodad: baseDoodads["bench_h"], pos: vec.I2{580, 80}},
		{BaseDoodad: baseDoodads["tree5"], pos: vec.I2{471, 141}},
		{BaseDoodad: baseDoodads["tree5"], pos: vec.I2{511, 245}},
		{BaseDoodad: baseDoodads["tree5"], pos: vec.I2{576, 145}},
		{BaseDoodad: baseDoodads["tree4"], pos: vec.I2{610, 250}},
		{BaseDoodad: baseDoodads["tree5"], pos: vec.I2{692, 148}},
		{BaseDoodad: baseDoodads["tree7"], pos: vec.I2{760, 243}},

		// Toward the south-west.
		{BaseDoodad: baseDoodads["tree1"], pos: vec.I2{78, 643}},
		{BaseDoodad: baseDoodads["tree2"], pos: vec.I2{99, 704}},
		{BaseDoodad: baseDoodads["tree1"], pos: vec.I2{158, 773}},
		{BaseDoodad: baseDoodads["tree2"], pos: vec.I2{121, 939}},
		{BaseDoodad: baseDoodads["tree4"], pos: vec.I2{183, 944}},

		// Trees in the east.
		{BaseDoodad: baseDoodads["tree1"], pos: vec.I2{932, 285}},
		{BaseDoodad: baseDoodads["tree2"], pos: vec.I2{934, 365}},
		{BaseDoodad: baseDoodads["tree8"], pos: vec.I2{885, 645}},
		{BaseDoodad: baseDoodads["tree2"], pos: vec.I2{937, 720}},
		{BaseDoodad: baseDoodads["tree7"], pos: vec.I2{765, 958}},
		{BaseDoodad: baseDoodads["tree2"], pos: vec.I2{669, 838}},
		{BaseDoodad: baseDoodads["tree4"], pos: vec.I2{732, 831}},

		// The W tree.
		//theW,
		{BaseDoodad: baseDoodads["tree1"], pos: vec.I2{860, 452}},

		// End-game in south-east.
		{BaseDoodad: baseDoodads["alamore"], pos: vec.I2{877, 877}},
	}
)

package lexis

import (
	"testing"
	"testing/fstest"
)

const testLdtkJSON = `{
	"defs": {
		"tilesets": [
			{"uid": 1, "relPath": "../../../../aseprite/floor.png", "tileGridSize": 32}
		]
	},
	"levels": [
		{
			"worldX": 0,
			"pxWid": 256,
			"pxHei": 224,
			"layerInstances": [
				{
					"__identifier": "Library_Floor",
					"__type": "Tiles",
					"__gridSize": 32,
					"__tilesetDefUid": 1,
					"gridTiles": [
						{"px": [0, 0], "src": [32, 0], "f": 0}
					],
					"entityInstances": []
				},
				{
					"__identifier": "Collisions",
					"__type": "Entities",
					"__gridSize": 32,
					"__tilesetDefUid": null,
					"gridTiles": [],
					"entityInstances": [
						{"__identifier": "WallCollider", "px": [0, 0], "width": 256, "height": 32}
					]
				}
			]
		}
	]
}`

func TestParseLdtkLevels(t *testing.T) {
	fsys := fstest.MapFS{
		"test.ldtk": &fstest.MapFile{Data: []byte(testLdtkJSON)},
	}
	project, err := parseLdtk(fsys, "test.ldtk")
	if err != nil {
		t.Fatalf("parseLdtk: %v", err)
	}
	if len(project.Levels) != 1 {
		t.Fatalf("levels = %d, want 1", len(project.Levels))
	}
	level := project.Levels[0]
	if level.PxWid != 256 || level.PxHei != 224 {
		t.Errorf("level size = %dx%d, want 256x224", level.PxWid, level.PxHei)
	}
}

func TestParseLdtkTilesets(t *testing.T) {
	fsys := fstest.MapFS{
		"test.ldtk": &fstest.MapFile{Data: []byte(testLdtkJSON)},
	}
	project, err := parseLdtk(fsys, "test.ldtk")
	if err != nil {
		t.Fatalf("parseLdtk: %v", err)
	}
	if len(project.Defs.Tilesets) != 1 {
		t.Fatalf("tilesets = %d, want 1", len(project.Defs.Tilesets))
	}
	ts := project.Defs.Tilesets[0]
	if ts.Uid != 1 {
		t.Errorf("tileset uid = %d, want 1", ts.Uid)
	}
	if ts.TileGridSize != 32 {
		t.Errorf("tileset gridSize = %d, want 32", ts.TileGridSize)
	}
}

func TestParseLdtkLayers(t *testing.T) {
	fsys := fstest.MapFS{
		"test.ldtk": &fstest.MapFile{Data: []byte(testLdtkJSON)},
	}
	project, err := parseLdtk(fsys, "test.ldtk")
	if err != nil {
		t.Fatalf("parseLdtk: %v", err)
	}
	layers := project.Levels[0].LayerInstances
	if len(layers) != 2 {
		t.Fatalf("layers = %d, want 2", len(layers))
	}

	floor := layers[0]
	if floor.Identifier != "Library_Floor" {
		t.Errorf("layer[0] identifier = %q, want Library_Floor", floor.Identifier)
	}
	if floor.GridSize != 32 {
		t.Errorf("layer[0] gridSize = %d, want 32", floor.GridSize)
	}
	if floor.TilesetDefUid == nil || *floor.TilesetDefUid != 1 {
		t.Errorf("layer[0] TilesetDefUid: got %v, want 1", floor.TilesetDefUid)
	}
	if len(floor.GridTiles) != 1 {
		t.Fatalf("layer[0] tiles = %d, want 1", len(floor.GridTiles))
	}
	tile := floor.GridTiles[0]
	if tile.Src != [2]int{32, 0} {
		t.Errorf("tile src = %v, want [32 0]", tile.Src)
	}

	collisions := layers[1]
	if collisions.TilesetDefUid != nil {
		t.Error("Collisions layer TilesetDefUid should be nil")
	}
	if len(collisions.EntityInstances) != 1 {
		t.Fatalf("collisions entities = %d, want 1", len(collisions.EntityInstances))
	}
	e := collisions.EntityInstances[0]
	if e.Width != 256 || e.Height != 32 {
		t.Errorf("entity size = %dx%d, want 256x32", e.Width, e.Height)
	}
}

func TestParseLdtkInvalidJSON(t *testing.T) {
	fsys := fstest.MapFS{
		"bad.ldtk": &fstest.MapFile{Data: []byte("not json {{{")},
	}
	_, err := parseLdtk(fsys, "bad.ldtk")
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestParseLdtkMissingFile(t *testing.T) {
	fsys := fstest.MapFS{}
	_, err := parseLdtk(fsys, "missing.ldtk")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

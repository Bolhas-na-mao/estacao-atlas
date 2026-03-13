package lexis

import (
	"embed"
	"encoding/json"
	"fmt"
)

type ldtkProject struct {
	Levels []ldtkLevel `json:"levels"`
}

type ldtkLevel struct {
	Identifier     string              `json:"identifier"`
	WorldX         int                 `json:"worldX"`
	WorldY         int                 `json:"worldY"`
	PxWid          int                 `json:"pxWid"`
	PxHei          int                 `json:"pxHei"`
	LayerInstances []ldtkLayerInstance `json:"layerInstances"`
}

type ldtkLayerInstance struct {
	Identifier      string       `json:"__identifier"`
	Type            string       `json:"__type"`
	GridSize        int          `json:"__gridSize"`
	GridTiles       []ldtkTile   `json:"gridTiles"`
	EntityInstances []ldtkEntity `json:"entityInstances"`
}

type ldtkEntity struct {
	Identifier string `json:"__identifier"`
	Px         [2]int `json:"px"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}

type ldtkTile struct {
	Px  [2]int `json:"px"`
	Src [2]int `json:"src"`
	F   int    `json:"f"`
}

func parseLdtk(fs embed.FS, path string) (*ldtkProject, error) {
	data, err := fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading ldtk: %w", err)
	}
	var p ldtkProject
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, fmt.Errorf("parsing ldtk: %w", err)
	}
	return &p, nil
}

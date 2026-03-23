package lexis

import (
	"embed"
	"image"
	"math"
	"path"
	"sort"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/logger"
	"github.com/Bolhas-na-mao/estacao-atlas/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

const tileSize = 32

type collider struct {
	x, y, w, h float64
}

type Room struct {
	width     int
	height    int
	colliders []collider
	img       *ebiten.Image
}

func newRoom(level ldtkLevel, sheets map[int]*ebiten.Image) *Room {
	r := &Room{
		width:  level.PxWid,
		height: level.PxHei,
	}
	img := ebiten.NewImage(level.PxWid, level.PxHei)

	for _, layer := range level.LayerInstances {
		if layer.Type != "Entities" {
			continue
		}
		for _, e := range layer.EntityInstances {
			if e.Identifier == "WallCollider" {
				r.colliders = append(r.colliders, collider{
					x: float64(e.Px[0]),
					y: float64(e.Px[1]),
					w: float64(e.Width),
					h: float64(e.Height),
				})
			}
		}
	}

	for i := len(level.LayerInstances) - 1; i >= 0; i-- {
		layer := level.LayerInstances[i]
		if layer.TilesetDefUid == nil {
			continue
		}
		sheet, ok := sheets[*layer.TilesetDefUid]
		if !ok {
			logger.Warn("tileset not loaded, skipping layer",
				"tilesetDefUid", *layer.TilesetDefUid,
				"layer", layer.Identifier,
				"layerIndex", i,
			)
			continue
		}
		renderTiles(img, layer.GridTiles, sheet, layer.GridSize)
	}

	r.img = img
	return r
}

func renderTiles(dst *ebiten.Image, tiles []ldtkTile, sheet *ebiten.Image, gridSize int) {
	for _, t := range tiles {
		sx, sy := t.Src[0], t.Src[1]
		src := sheet.SubImage(image.Rect(sx, sy, sx+gridSize, sy+gridSize)).(*ebiten.Image)

		op := &ebiten.DrawImageOptions{}
		if t.F&1 != 0 {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(float64(gridSize), 0)
		}
		if t.F&2 != 0 {
			op.GeoM.Scale(1, -1)
			op.GeoM.Translate(0, float64(gridSize))
		}
		op.GeoM.Translate(float64(t.Px[0]), float64(t.Px[1]))
		dst.DrawImage(src, op)
	}
}

func (r *Room) isSolid(x, y float64) bool {
	for _, c := range r.colliders {
		if x >= c.x && x < c.x+c.w && y >= c.y && y < c.y+c.h {
			return true
		}
	}
	return false
}

type WorldMap struct {
	rooms      []*Room
	currentIdx int
}

func newWorldMap(project *ldtkProject, fs embed.FS) *WorldMap {
	sheets := make(map[int]*ebiten.Image)
	for _, ts := range project.Defs.Tilesets {
		filename := path.Base(ts.RelPath)
		img, err := ui.RenderAsset(fs, "assets/tilesets/"+filename)
		if err != nil {
			logger.Fatal("loading tileset failed", "file", filename, "err", err)
		}
		sheets[ts.Uid] = img
	}

	logger.Info("world map loading", "tilesets", len(sheets))

	levels := make([]ldtkLevel, len(project.Levels))
	copy(levels, project.Levels)
	sort.Slice(levels, func(i, j int) bool {
		return levels[i].WorldX < levels[j].WorldX
	})

	rooms := make([]*Room, len(levels))
	for i, level := range levels {
		rooms[i] = newRoom(level, sheets)
	}

	logger.Info("world map ready", "rooms", len(rooms))
	return &WorldMap{rooms: rooms}
}

func (w *WorldMap) current() *Room {
	return w.rooms[w.currentIdx]
}

func (w *WorldMap) draw(screen *ebiten.Image, cam *Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(heroScale, heroScale)
	op.GeoM.Translate(-cam.x*heroScale, -cam.y*heroScale)
	screen.DrawImage(w.current().img, op)
}

func clamp(v, lo, hi float64) float64 {
	return math.Max(lo, math.Min(v, hi))
}

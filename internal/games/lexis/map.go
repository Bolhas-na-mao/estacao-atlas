package lexis

import (
	"image"
	"math"
	"sort"

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

func newRoom(level ldtkLevel, wallSheet, floorSheet, bookshelfSheet *ebiten.Image) *Room {
	r := &Room{
		width:  level.PxWid,
		height: level.PxHei,
	}
	img := ebiten.NewImage(level.PxWid, level.PxHei)

	var wallTiles, bookshelfTiles []ldtkTile
	for _, layer := range level.LayerInstances {
		switch layer.Identifier {
		case "Library_Floor":
			renderTiles(img, layer.GridTiles, floorSheet, 0)
		case "Library_Wall":
			wallTiles = layer.GridTiles
		case "Bookshelf":
			bookshelfTiles = layer.GridTiles
		case "Collisions":
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
	}
	renderTiles(img, wallTiles, wallSheet, 0)
	renderTiles(img, bookshelfTiles, bookshelfSheet, -tileSize)

	r.img = img
	return r
}

func renderTiles(dst *ebiten.Image, tiles []ldtkTile, sheet *ebiten.Image, yOffset int) {
	for _, t := range tiles {
		sx, sy := t.Src[0], t.Src[1]
		src := sheet.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)

		op := &ebiten.DrawImageOptions{}
		if t.F&1 != 0 {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(tileSize, 0)
		}
		if t.F&2 != 0 {
			op.GeoM.Scale(1, -1)
			op.GeoM.Translate(0, tileSize)
		}
		op.GeoM.Translate(float64(t.Px[0]), float64(t.Px[1]+yOffset))
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

func newWorldMap(project *ldtkProject, wallSheet, floorSheet, bookshelfSheet *ebiten.Image) *WorldMap {
	levels := make([]ldtkLevel, len(project.Levels))
	copy(levels, project.Levels)
	sort.Slice(levels, func(i, j int) bool {
		return levels[i].WorldX < levels[j].WorldX
	})

	rooms := make([]*Room, len(levels))
	for i, level := range levels {
		rooms[i] = newRoom(level, wallSheet, floorSheet, bookshelfSheet)
	}
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

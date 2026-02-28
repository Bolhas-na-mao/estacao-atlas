package lexis

import "github.com/hajimehoshi/ebiten/v2"

const (
	roomWidth   = 256
	roomHeight  = 112
	worldWidth  = roomWidth * 2
	worldHeight = roomHeight
)

type Room struct {
	sprite *ebiten.Image
	npcs   []Npc
	worldX float64
}

type Map struct {
	rooms []Room
}

func (m *Map) nearestNpc(playerX, maxDist float64) *Npc {
	for i := range m.rooms {
		for j := range m.rooms[i].npcs {
			npc := &m.rooms[i].npcs[j]
			dx := npc.x - playerX
			if dx < 0 {
				dx = -dx
			}
			if dx < maxDist && len(npc.dialogue) > 0 {
				return npc
			}
		}
	}
	return nil
}

func (m *Map) draw(screen *ebiten.Image, cam *Camera) {
	for i := range m.rooms {
		room := &m.rooms[i]
		sx, sy := cam.toScreen(room.worldX, 0)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(heroScale, heroScale)
		op.GeoM.Translate(sx, sy)
		screen.DrawImage(room.sprite, op)

		for j := range room.npcs {
			npc := &room.npcs[j]
			nx, ny := cam.toScreen(npc.x, npc.y)
			nop := &ebiten.DrawImageOptions{}
			nop.GeoM.Scale(heroScale, heroScale)
			nop.GeoM.Translate(nx, ny)
			screen.DrawImage(npc.currentImage(), nop)
		}
	}
}

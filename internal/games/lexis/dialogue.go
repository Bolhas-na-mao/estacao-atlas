package lexis

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	dialogueBoxMargin = 40
	dialogueBoxY      = 575
	dialogueBoxHeight = 120
	dialoguePadding   = 16
	interactionRange  = 120.0
)

func fillRect(screen *ebiten.Image, x, y, w, h int, clr color.Color) {
	pixel := ebiten.NewImage(1, 1)
	pixel.Fill(clr)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(w), float64(h))
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(pixel, op)
}

type DialogueState struct {
	active bool
	npc    *Npc
	line   int
}

func (d *DialogueState) start(npc *Npc) {
	d.active = true
	d.npc = npc
	d.line = 0
}

func (d *DialogueState) advance() {
	if !d.active {
		return
	}
	d.line++
	if d.line >= len(d.npc.dialogue) {
		d.active = false
		d.npc = nil
		d.line = 0
	}
}

func (d *DialogueState) draw(screen *ebiten.Image) {
	if !d.active {
		return
	}

	const (
		x = dialogueBoxMargin
		y = dialogueBoxY
		w = screenWidth - dialogueBoxMargin*2
		h = dialogueBoxHeight
	)

	// Border
	fillRect(screen, x-2, y-2, w+4, h+4, color.RGBA{180, 160, 255, 200})

	// Background
	fillRect(screen, x, y, w, h, color.RGBA{15, 10, 30, 230})

	// NPC name
	name := fmt.Sprintf("[ %s ]", d.npc.name)
	ebitenutil.DebugPrintAt(screen, name, int(x)+dialoguePadding, int(y)+dialoguePadding)

	// Dialogue text
	if d.line < len(d.npc.dialogue) {
		ebitenutil.DebugPrintAt(screen, d.npc.dialogue[d.line], int(x)+dialoguePadding, int(y)+dialoguePadding+20)
	}

	continueHint := "[ ESPAÇO ]"
	ebitenutil.DebugPrintAt(screen, continueHint, int(x+w)-80, int(y+h)-18)
}

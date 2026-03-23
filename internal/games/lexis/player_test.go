package lexis

import (
	"math"
	"testing"
)

func TestPlayerHitsSolidNoCollision(t *testing.T) {
	p := &Player{x: 100, y: 100}
	if p.hitsSolid(100, 100, func(x, y float64) bool { return false }) {
		t.Error("expected no collision with empty world")
	}
}

func TestPlayerHitsSolidAllSolid(t *testing.T) {
	p := &Player{x: 100, y: 100}
	if !p.hitsSolid(100, 100, func(x, y float64) bool { return true }) {
		t.Error("expected collision with all-solid world")
	}
}

func TestPlayerHitsSolidBottomHalfOnly(t *testing.T) {
	p := &Player{x: 100, y: 100}
	// spriteSize=32; hitbox top = y + 16 = 116
	solidAtOrBelowMid := func(x, y float64) bool { return y >= 116 }
	if !p.hitsSolid(100, 100, solidAtOrBelowMid) {
		t.Error("expected collision when solid at lower half")
	}
	// Solid only above the hitbox — no collision
	solidAbove := func(x, y float64) bool { return y < 100 }
	if p.hitsSolid(100, 100, solidAbove) {
		t.Error("expected no collision when solid only above hitbox")
	}
}

func TestPlayerMoveNoInput(t *testing.T) {
	p := &Player{x: 100, y: 100}
	p.move(0, 0, func(x, y float64) bool { return false })
	if p.x != 100 || p.y != 100 {
		t.Errorf("position changed on zero input: (%.2f, %.2f)", p.x, p.y)
	}
	if p.isWalking {
		t.Error("isWalking should be false on zero input")
	}
}

func TestPlayerMoveRight(t *testing.T) {
	p := &Player{x: 100, y: 100}
	p.move(1, 0, func(x, y float64) bool { return false })
	if p.x <= 100 {
		t.Errorf("expected x to increase, got %.2f", p.x)
	}
	if p.currDir != East {
		t.Errorf("expected direction East, got %v", p.currDir)
	}
	if !p.isWalking {
		t.Error("expected isWalking=true after moving")
	}
}

func TestPlayerMoveLeft(t *testing.T) {
	p := &Player{x: 100, y: 100}
	p.move(-1, 0, func(x, y float64) bool { return false })
	if p.x >= 100 {
		t.Errorf("expected x to decrease, got %.2f", p.x)
	}
	if p.currDir != West {
		t.Errorf("expected direction West, got %v", p.currDir)
	}
}

func TestPlayerMoveBlocked(t *testing.T) {
	p := &Player{x: 100, y: 100}
	p.move(1, 0, func(x, y float64) bool { return true })
	if p.x != 100 {
		t.Errorf("expected x unchanged when blocked, got %.2f", p.x)
	}
}

func TestPlayerMoveDiagonalNormalized(t *testing.T) {
	// Diagonal movement should cover same total distance as cardinal
	diag := &Player{x: 100, y: 100}
	diag.move(1, 1, func(x, y float64) bool { return false })
	cardinal := &Player{x: 100, y: 100}
	cardinal.move(1, 0, func(x, y float64) bool { return false })

	// Each axis of diagonal should be less than cardinal (normalized)
	if diag.x-100 >= cardinal.x-100 {
		t.Errorf("diagonal x-component (%.2f) should be less than cardinal (%.2f)",
			diag.x-100, cardinal.x-100)
	}
	diagDist := math.Hypot(diag.x-100, diag.y-100)
	cardDist := math.Hypot(cardinal.x-100, cardinal.y-100)
	if math.Abs(diagDist-cardDist) > 1e-6 {
		t.Errorf("diagonal distance (%.6f) should equal cardinal distance (%.6f)", diagDist, cardDist)
	}
}

func TestPlayerDirectionUpdates(t *testing.T) {
	tests := []struct {
		dx, dy float64
		dir    Direction
	}{
		{1, 0, East},
		{-1, 0, West},
		{0, 1, South},
		{0, -1, North},
	}
	for _, tt := range tests {
		p := &Player{x: 100, y: 100}
		p.move(tt.dx, tt.dy, func(x, y float64) bool { return false })
		if p.currDir != tt.dir {
			t.Errorf("move(%.0f,%.0f) direction=%v, want %v", tt.dx, tt.dy, p.currDir, tt.dir)
		}
	}
}

func TestPlayerUpdateAnimation(t *testing.T) {
	p := &Player{isWalking: true}
	for i := 0; i < animationSpeed; i++ {
		p.update()
	}
	if p.animFrame != 1 {
		t.Errorf("animFrame = %d after %d ticks, want 1", p.animFrame, animationSpeed)
	}
}

func TestPlayerUpdateAnimationWraps(t *testing.T) {
	p := &Player{isWalking: true}
	for i := 0; i < animationSpeed*walkFrames; i++ {
		p.update()
	}
	if p.animFrame != 0 {
		t.Errorf("animFrame = %d after full cycle, want 0", p.animFrame)
	}
}

func TestPlayerUpdateStopsAnimation(t *testing.T) {
	p := &Player{isWalking: false, animFrame: 2, animTick: 5}
	p.update()
	if p.animFrame != 0 || p.animTick != 0 {
		t.Errorf("expected frame=0 tick=0 when not walking, got frame=%d tick=%d", p.animFrame, p.animTick)
	}
}

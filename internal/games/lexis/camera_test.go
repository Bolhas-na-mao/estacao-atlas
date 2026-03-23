package lexis

import "testing"

func TestClamp(t *testing.T) {
	tests := []struct{ v, lo, hi, want float64 }{
		{5, 0, 10, 5},
		{-5, 0, 10, 0},
		{15, 0, 10, 10},
		{0, 0, 10, 0},
		{10, 0, 10, 10},
	}
	for _, tt := range tests {
		if got := clamp(tt.v, tt.lo, tt.hi); got != tt.want {
			t.Errorf("clamp(%v, %v, %v) = %v, want %v", tt.v, tt.lo, tt.hi, got, tt.want)
		}
	}
}

func TestCameraViewDimensions(t *testing.T) {
	cam := &Camera{sw: 1280, sh: 720}
	if got := cam.viewW(); got != 512 {
		t.Errorf("viewW() = %v, want 512", got)
	}
	if got := cam.viewH(); got != 288 {
		t.Errorf("viewH() = %v, want 288", got)
	}
}

func TestCameraSnapToCenter(t *testing.T) {
	cam := &Camera{sw: 1280, sh: 720}
	cam.snapTo(1000, 500, 2000, 2000)
	if cam.x != 744 {
		t.Errorf("cam.x = %v, want 744", cam.x)
	}
	if cam.y != 356 {
		t.Errorf("cam.y = %v, want 356", cam.y)
	}
}

func TestCameraSnapToClampsMin(t *testing.T) {
	cam := &Camera{sw: 1280, sh: 720}
	cam.snapTo(0, 0, 2000, 2000)
	if cam.x != 0 {
		t.Errorf("cam.x = %v, want 0", cam.x)
	}
	if cam.y != 0 {
		t.Errorf("cam.y = %v, want 0", cam.y)
	}
}

func TestCameraSnapToClampMax(t *testing.T) {
	cam := &Camera{sw: 1280, sh: 720}
	cam.snapTo(9999, 9999, 800, 600)
	if cam.x != 288 {
		t.Errorf("cam.x = %v, want 288", cam.x)
	}
	if cam.y != 312 {
		t.Errorf("cam.y = %v, want 312", cam.y)
	}
}

func TestCameraUpdateDeadZoneNoop(t *testing.T) {
	cam := &Camera{sw: 1280, sh: 720}
	cam.snapTo(500, 500, 2000, 2000)
	heroCenterX := cam.x + cam.viewW()/2
	heroCenterY := cam.y + cam.viewH()/2
	prevTargetX := cam.targetX
	prevTargetY := cam.targetY
	cam.update(heroCenterX, heroCenterY, 2000, 2000)
	if cam.targetX != prevTargetX || cam.targetY != prevTargetY {
		t.Error("camera target changed when hero was at viewport center (inside dead zone)")
	}
}

func TestCameraUpdateMovesTargetRight(t *testing.T) {
	cam := &Camera{sw: 1280, sh: 720}
	cam.snapTo(500, 500, 2000, 2000)
	prevTargetX := cam.targetX
	heroCenterX := cam.x + cam.viewW()/2 + deadZoneW + 50
	heroCenterY := cam.y + cam.viewH()/2
	cam.update(heroCenterX, heroCenterY, 2000, 2000)
	if cam.targetX <= prevTargetX {
		t.Errorf("targetX did not increase: got %v, was %v", cam.targetX, prevTargetX)
	}
}

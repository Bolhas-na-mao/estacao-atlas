package lexis

import "testing"

func TestIsSolidInsideCollider(t *testing.T) {
	r := &Room{
		colliders: []collider{{x: 0, y: 0, w: 100, h: 50}},
	}
	if !r.isSolid(50, 25) {
		t.Error("expected (50,25) to be solid")
	}
}

func TestIsSolidOutsideCollider(t *testing.T) {
	r := &Room{
		colliders: []collider{{x: 0, y: 0, w: 100, h: 50}},
	}
	if r.isSolid(150, 25) {
		t.Error("expected (150,25) to not be solid")
	}
}

func TestIsSolidExclusiveBoundary(t *testing.T) {
	r := &Room{
		colliders: []collider{{x: 0, y: 0, w: 100, h: 50}},
	}
	// right edge is exclusive
	if r.isSolid(100, 25) {
		t.Error("x=100 (exclusive boundary) should not be solid")
	}
	// bottom edge is exclusive
	if r.isSolid(50, 50) {
		t.Error("y=50 (exclusive boundary) should not be solid")
	}
}

func TestIsSolidInclusiveBoundary(t *testing.T) {
	r := &Room{
		colliders: []collider{{x: 10, y: 10, w: 50, h: 50}},
	}
	if !r.isSolid(10, 10) {
		t.Error("origin (10,10) should be solid (inclusive)")
	}
}

func TestIsSolidMultipleColliders(t *testing.T) {
	r := &Room{
		colliders: []collider{
			{x: 0, y: 0, w: 50, h: 50},
			{x: 100, y: 100, w: 50, h: 50},
		},
	}
	if !r.isSolid(25, 25) {
		t.Error("expected (25,25) to be solid (first collider)")
	}
	if !r.isSolid(125, 125) {
		t.Error("expected (125,125) to be solid (second collider)")
	}
	if r.isSolid(75, 75) {
		t.Error("expected (75,75) to not be solid (gap between colliders)")
	}
}

func TestIsSolidNoColliders(t *testing.T) {
	r := &Room{}
	if r.isSolid(50, 50) {
		t.Error("expected no solid when no colliders exist")
	}
}

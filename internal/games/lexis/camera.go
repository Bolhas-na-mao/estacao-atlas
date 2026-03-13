package lexis

import "math"

const (
	screenWidth  = 1280
	screenHeight = 720

	deadZoneW = 50.0
	deadZoneH = 40.0

	cameraLerp = 0.1
)

type Camera struct {
	x, y    float64
	targetX float64
	targetY float64
}

func viewW() float64 { return float64(screenWidth) / heroScale }

func viewH() float64 { return float64(screenHeight) / heroScale }

func newCamera(heroCenterX, heroCenterY float64, roomW, roomH int) *Camera {
	c := &Camera{}
	c.snapTo(heroCenterX, heroCenterY, roomW, roomH)
	return c
}

func (c *Camera) update(heroCenterX, heroCenterY float64, roomW, roomH int) {
	vw := viewW()
	vh := viewH()

	heroViewX := heroCenterX - c.x
	heroViewY := heroCenterY - c.y

	if heroViewX < vw/2-deadZoneW {
		c.targetX = heroCenterX - (vw/2 - deadZoneW)
	} else if heroViewX > vw/2+deadZoneW {
		c.targetX = heroCenterX - (vw/2 + deadZoneW)
	}
	if heroViewY < vh/2-deadZoneH {
		c.targetY = heroCenterY - (vh/2 - deadZoneH)
	} else if heroViewY > vh/2+deadZoneH {
		c.targetY = heroCenterY - (vh/2 + deadZoneH)
	}

	maxX := math.Max(0, float64(roomW)-vw)
	maxY := math.Max(0, float64(roomH)-vh)
	c.targetX = math.Max(0, math.Min(c.targetX, maxX))
	c.targetY = math.Max(0, math.Min(c.targetY, maxY))

	c.x += (c.targetX - c.x) * cameraLerp
	c.y += (c.targetY - c.y) * cameraLerp
}

// immediately centers the camera on the hero and clamps to the room, with no lerp (for room transitions.)
func (c *Camera) snapTo(heroCenterX, heroCenterY float64, roomW, roomH int) {
	vw := viewW()
	vh := viewH()
	maxX := math.Max(0, float64(roomW)-vw)
	maxY := math.Max(0, float64(roomH)-vh)
	c.x = math.Max(0, math.Min(heroCenterX-vw/2, maxX))
	c.y = math.Max(0, math.Min(heroCenterY-vh/2, maxY))
	c.targetX = c.x
	c.targetY = c.y
}

package lexis

import "math"

const (
	deadZoneW  = 50.0
	deadZoneH  = 40.0
	cameraLerp = 0.1
)

type Camera struct {
	x, y    float64
	targetX float64
	targetY float64
	sw, sh  int
}

func (c *Camera) viewW() float64 { return float64(c.sw) / heroScale }
func (c *Camera) viewH() float64 { return float64(c.sh) / heroScale }

func newCamera(heroCenterX, heroCenterY float64, roomW, roomH, sw, sh int) *Camera {
	c := &Camera{sw: sw, sh: sh}
	c.snapTo(heroCenterX, heroCenterY, roomW, roomH)
	return c
}

func (c *Camera) update(heroCenterX, heroCenterY float64, roomW, roomH int) {
	vw := c.viewW()
	vh := c.viewH()

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

func (c *Camera) snapTo(heroCenterX, heroCenterY float64, roomW, roomH int) {
	vw := c.viewW()
	vh := c.viewH()
	maxX := math.Max(0, float64(roomW)-vw)
	maxY := math.Max(0, float64(roomH)-vh)
	c.x = math.Max(0, math.Min(heroCenterX-vw/2, maxX))
	c.y = math.Max(0, math.Min(heroCenterY-vh/2, maxY))
	c.targetX = c.x
	c.targetY = c.y
}

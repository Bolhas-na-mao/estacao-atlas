package lexis

const (
	screenWidth      = 1280
	screenHeight     = 720
	deadZoneHalfW    = 80.0
	cameraLerpFactor = 0.1
	verticalOffset   = (screenHeight - roomHeight*heroScale) / 2
)

type Camera struct {
	x float64
}

func (c *Camera) update(playerX, worldWidth float64) {
	playerScreenX := (playerX - c.x) * heroScale

	dzLeft := float64(screenWidth)/2 - deadZoneHalfW
	dzRight := float64(screenWidth)/2 + deadZoneHalfW

	targetX := c.x
	if playerScreenX < dzLeft {
		targetX = playerX - dzLeft/heroScale
	} else if playerScreenX > dzRight {
		targetX = playerX - dzRight/heroScale
	}

	c.x += (targetX - c.x) * cameraLerpFactor

	maxCamX := worldWidth - float64(screenWidth)/heroScale
	if c.x < 0 {
		c.x = 0
	} else if maxCamX > 0 && c.x > maxCamX {
		c.x = maxCamX
	}
}

func (c *Camera) toScreen(worldX, worldY float64) (float64, float64) {
	return (worldX - c.x) * heroScale, worldY*heroScale + verticalOffset
}

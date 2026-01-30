package launcher

import (
	"image/color"
	"log"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type LauncherState int

const (
	StateMenu LauncherState = iota
	StatePlaying
)

type Launcher struct {
	img          *ebiten.Image
	state        LauncherState
	screenWidth  int
	screenHeight int
}

const SCREEN_WIDTH = 1280
const SCREEN_HEIGHT = 720
const CELL_SIZE = 30
const PADDING = 10
const LOGO_PATH = "assets/atlas_logo.png"
const LAUNCHER_TITLE = "Estação Atlas"

func (l *Launcher) Update() error {
	return nil
}

func (l *Launcher) drawMenu(screen *ebiten.Image) {
	screen.Fill(color.RGBA{228, 228, 228, 255})

	ui.DrawGrid(color.RGBA{217, 218, 224, 255}, screen, l.screenWidth, l.screenHeight, PADDING, CELL_SIZE)

	if l.img != nil {
		w, _ := l.img.Bounds().Dx(), l.img.Bounds().Dy()

		op := &ebiten.DrawImageOptions{}

		scale := 0.35
		op.GeoM.Scale(scale, scale)

		x := (float64(l.screenWidth) - (float64(w) * scale)) / 2
		y := float64(l.screenHeight) / 100

		op.GeoM.Translate(x, y)

		screen.DrawImage(l.img, op)
	}
}

func (l *Launcher) Draw(screen *ebiten.Image) {
	switch l.state {
	case StateMenu:
		l.drawMenu(screen)
	default:
		l.drawMenu(screen)
	}
}

func (l *Launcher) GetArea() (int, int) {
	return l.screenWidth, l.screenHeight
}

func (l *Launcher) Layout(outsideWidth, outsideHeight int) (int, int) {
	return l.GetArea()
}

func NewLauncher() *Launcher {
	img, _, err := ebitenutil.NewImageFromFile(LOGO_PATH)
	if err != nil {
		log.Fatal(err)
	}

	return &Launcher{
		img:          img,
		state:        StateMenu,
		screenWidth:  SCREEN_WIDTH,
		screenHeight: SCREEN_HEIGHT,
	}
}

func RunLauncher() {
	launcher := NewLauncher()

	width, height := launcher.GetArea()

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(LAUNCHER_TITLE)

	if err := ebiten.RunGame(launcher); err != nil {
		log.Fatal(err)
	}
}

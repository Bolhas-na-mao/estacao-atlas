package launcher

import (
	"image/color"
	"log"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/games"
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
	gameButtons  []*ui.Button
}

const SCREEN_WIDTH = 1280
const SCREEN_HEIGHT = 720
const CELL_SIZE = 30
const PADDING = 10
const LOGO_PATH = "internal/launcher/assets/atlas_logo.png"
const LAUNCHER_TITLE = "Estação Atlas"

func (l *Launcher) Update() error {

	if l.state == StatePlaying {
		return games.UpdateCurrentGame()
	}

	if l.state == StateMenu {
		availableGames := games.ListGames()

		for i, btn := range l.gameButtons {
			if btn.Update() {
				selectedGame := availableGames[i]

				games.SetCurrentGame(selectedGame.ID)

				l.state = StatePlaying
			}
		}
	}

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

	for _, btn := range l.gameButtons {
		btn.Draw(screen)
	}
}

func (l *Launcher) Draw(screen *ebiten.Image) {
	switch l.state {
	case StateMenu:
		l.drawMenu(screen)
	case StatePlaying:
		l.Play(screen)

	default:
		l.drawMenu(screen)
	}
}

func (l *Launcher) Play(screen *ebiten.Image) {

	err := games.PlayCurrentGame(screen)

	if err != nil {
		l.drawMenu(screen)
	}
}

func (l *Launcher) GetArea() (int, int) {
	return l.screenWidth, l.screenHeight
}

func (l *Launcher) Layout(outsideWidth, outsideHeight int) (int, int) {
	return l.GetArea()
}

func (l *Launcher) RenderButtons() {
	availableGames := games.ListGames()
	startY := 300

	for i, g := range availableGames {
		btn := ui.NewButton(
			(SCREEN_WIDTH-200)/2,
			startY+(i*60),
			200, 50,
			g.Name,
		)

		l.gameButtons = append(l.gameButtons, btn)
	}

}

func NewLauncher() *Launcher {
	img, _, err := ebitenutil.NewImageFromFile(LOGO_PATH)
	if err != nil {
		log.Fatal(err)
	}

	l := &Launcher{
		img:          img,
		state:        StateMenu,
		screenWidth:  SCREEN_WIDTH,
		screenHeight: SCREEN_HEIGHT,
		gameButtons:  []*ui.Button{},
	}

	l.RenderButtons()

	return l
}

func RunLauncher() {
	launcher := NewLauncher()

	width, height := launcher.GetArea()

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(LAUNCHER_TITLE)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(launcher); err != nil {
		log.Fatal(err)
	}
}

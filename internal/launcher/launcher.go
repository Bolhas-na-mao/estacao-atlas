package launcher

import (
	"embed"
	"image/color"
	_ "image/png"
	"log"

	"github.com/Bolhas-na-mao/estacao-atlas/internal/games"
	_ "github.com/Bolhas-na-mao/estacao-atlas/internal/games/lexis"
	"github.com/Bolhas-na-mao/estacao-atlas/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//go:embed assets/*
var assets embed.FS

type LauncherState int

const (
	StateMenu LauncherState = iota
	StatePlaying
)

const (
	defaultScreenWidth  = 1280
	defaultScreenHeight = 720
	cellSize            = 30
	gridPadding         = 10
	logoPath            = "assets/atlas_logo.png"
	launcherTitle       = "Estação Atlas"
	logoScale           = 0.35
	buttonStartY        = 300
	buttonWidth         = 200
	buttonHeight        = 50
	buttonSpacing       = 60
)

type Launcher struct {
	img          *ebiten.Image
	state        LauncherState
	screenWidth  int
	screenHeight int
	gameButtons  []*ui.Button
	gameInfos    []games.GameInfo
	currentGame  games.Game
}

func (l *Launcher) Update() error {
	if l.state == StatePlaying {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			l.currentGame = nil
			l.state = StateMenu
			return nil
		}
		return l.currentGame.Update()
	}

	for i, btn := range l.gameButtons {
		if btn.Update() {
			l.currentGame = l.gameInfos[i].New()
			l.state = StatePlaying
		}
	}

	return nil
}

func (l *Launcher) drawMenu(screen *ebiten.Image) {
	screen.Fill(color.RGBA{228, 228, 228, 255})

	ui.DrawGrid(color.RGBA{217, 218, 224, 255}, screen, l.screenWidth, l.screenHeight, gridPadding, cellSize)

	if l.img != nil {
		w := l.img.Bounds().Dx()

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(logoScale, logoScale)

		x := (float64(l.screenWidth) - (float64(w) * logoScale)) / 2
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
	case StatePlaying:
		l.currentGame.Draw(screen)
	default:
		l.drawMenu(screen)
	}
}

func (l *Launcher) Layout(outsideWidth, outsideHeight int) (int, int) {
	return l.screenWidth, l.screenHeight
}

func (l *Launcher) renderButtons() {
	for i, g := range l.gameInfos {
		btn := ui.NewButton(
			(defaultScreenWidth-buttonWidth)/2,
			buttonStartY+(i*buttonSpacing),
			buttonWidth, buttonHeight,
			g.Name,
		)
		l.gameButtons = append(l.gameButtons, btn)
	}
}

func NewLauncher() (*Launcher, error) {
	img, err := ui.RenderAsset(assets, logoPath)
	if err != nil {
		return nil, err
	}

	launcher := &Launcher{
		img:          img,
		state:        StateMenu,
		screenWidth:  defaultScreenWidth,
		screenHeight: defaultScreenHeight,
		gameInfos:    games.ListGames(),
	}

	launcher.renderButtons()

	return launcher, nil
}

func RunLauncher() {
	launcher, err := NewLauncher()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(launcher.screenWidth, launcher.screenHeight)
	ebiten.SetWindowTitle(launcherTitle)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(launcher); err != nil {
		log.Fatal(err)
	}
}

package ui

import (
	"bytes"
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func RenderAsset(assets embed.FS, path string) (*ebiten.Image, error) {
	logoData, err := assets.ReadFile(path)
	if err != nil {
		return nil, err
	}

	logoImg, _, err := image.Decode(bytes.NewReader(logoData))

	return ebiten.NewImageFromImage(logoImg), nil
}

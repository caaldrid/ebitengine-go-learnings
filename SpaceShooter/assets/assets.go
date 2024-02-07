package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed Backgrounds/* PNG/*
var assetsFS embed.FS

type Assets struct {
	Background *ebiten.Image
	Player     *ebiten.Image
}

func (a *Assets) loadImageFromFS(path string) *ebiten.Image {
	f, err := assetsFS.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func NewAssets() *Assets {
	a := &Assets{}
	a.Background = a.loadImageFromFS("Backgrounds/darkPurple.png")
	a.Player = a.loadImageFromFS("PNG/playerShip1_orange.png")

	return a
}

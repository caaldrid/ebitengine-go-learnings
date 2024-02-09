package assets

import (
	"embed"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed Backgrounds/* PNG/*
var assetsFS embed.FS

type Assets struct {
	Background *ebiten.Image
	Player     *ebiten.Image
	Meteors    []*ebiten.Image
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

func (a *Assets) loadMultipleImagesFromFS(globPath string) []*ebiten.Image {
	matches, err := fs.Glob(assetsFS, globPath)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = a.loadImageFromFS(match)
	}

	return images
}

func NewAssets() *Assets {
	a := &Assets{}
	a.Background = a.loadImageFromFS("Backgrounds/darkPurple.png")
	a.Player = a.loadImageFromFS("PNG/playerShip1_orange.png")
	a.Meteors = a.loadMultipleImagesFromFS("PNG/Meteors/*.png")

	return a
}

func CalcCenter(sprite *ebiten.Image) (float64, float64) {
	// Find the center of the sprite
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	return halfW, halfH
}

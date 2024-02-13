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

type Asset struct {
	Sprite *ebiten.Image
}

func (a *Asset) CalcCenter() (float64, float64) {
	// Find the center of the sprite
	bounds := a.Sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	return halfW, halfH
}

type MissileAssets struct {
	Bullet    *Asset
	Explosion *Asset
}

type BulletAssets struct {
	Basic   *Asset
	Upgrade *Asset
	Missile *MissileAssets
}

type Assets struct {
	Background *Asset
	Player     *Asset
	Meteors    []*Asset
	Bullets    *BulletAssets
}

func (as *Assets) loadImageFromFS(path string) *Asset {
	f, err := assetsFS.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return &Asset{
		Sprite: ebiten.NewImageFromImage(img),
	}
}

func (as *Assets) loadMultipleImagesFromFS(globPath string) []*Asset {
	matches, err := fs.Glob(assetsFS, globPath)
	if err != nil {
		panic(err)
	}

	images := make([]*Asset, len(matches))
	for i, match := range matches {
		images[i] = as.loadImageFromFS(match)
	}

	return images
}

func NewAssets() *Assets {
	as := &Assets{}
	as.Background = as.loadImageFromFS("Backgrounds/darkPurple.png")
	as.Player = as.loadImageFromFS("PNG/playerShip1_orange.png")
	as.Meteors = as.loadMultipleImagesFromFS("PNG/Meteors/*.png")
	as.Bullets = &BulletAssets{
		Basic:   as.loadImageFromFS("PNG/Lasers/laserBlue.png"),
		Upgrade: as.loadImageFromFS("PNG/Lasers/Upgrades/laserGreenUpgrade.png"),
		Missile: &MissileAssets{
			Bullet:    as.loadImageFromFS("PNG/Lasers/Upgrades/laserRedMissile.png"),
			Explosion: as.loadImageFromFS("PNG/Lasers/Upgrades/laserRedMissileExplosion.png"),
		},
	}

	return as
}

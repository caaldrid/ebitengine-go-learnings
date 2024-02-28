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

func loadImageFromFS(path string) *ebiten.Image {
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

func loadMultipleImagesFromFS(globPath string) []*ebiten.Image {
	matches, err := fs.Glob(assetsFS, globPath)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = loadImageFromFS(match)
	}

	return images
}

func LoadSprites() *Sprites {
	s := &Sprites{}
	s.Background = loadImageFromFS("Backgrounds/darkPurple.png")
	s.Player = loadImageFromFS("PNG/playerShip1_orange.png")
	s.Meteors = loadMultipleImagesFromFS("PNG/Meteors/*.png")
	s.Bullets = &BulletSprites{
		Basic: &MissileSprites{
			Bullet:    loadImageFromFS("PNG/Lasers/laserBlue.png"),
			Explosion: loadImageFromFS("PNG/Lasers/laserBlueExplosion.png"),
		},
		Upgrade: &MissileSprites{
			Bullet:    loadImageFromFS("PNG/Lasers/Upgrades/laserGreenUpgrade.png"),
			Explosion: loadImageFromFS("PNG/Lasers/Upgrades/laserGreenUpgradeExplosion.png"),
		},
		Missile: &MissileSprites{
			Bullet:    loadImageFromFS("PNG/Lasers/Upgrades/laserRedMissile.png"),
			Explosion: loadImageFromFS("PNG/Lasers/Upgrades/laserRedMissileExplosion.png"),
		},
	}

	return s
}

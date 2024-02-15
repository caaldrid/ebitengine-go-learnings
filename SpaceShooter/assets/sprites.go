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

type MissileSprites struct {
	Bullet    *ebiten.Image
	Explosion *ebiten.Image
}

type BulletSprites struct {
	Basic   *MissileSprites
	Upgrade *MissileSprites
	Missile *MissileSprites
}

type Sprites struct {
	Background *ebiten.Image
	Player     *ebiten.Image
	Meteors    []*ebiten.Image
	Bullets    *BulletSprites
}

func (s *Sprites) loadImageFromFS(path string) *ebiten.Image {
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

func (s *Sprites) loadMultipleImagesFromFS(globPath string) []*ebiten.Image {
	matches, err := fs.Glob(assetsFS, globPath)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = s.loadImageFromFS(match)
	}

	return images
}

func LoadSprites() *Sprites {
	s := &Sprites{}
	s.Background = s.loadImageFromFS("Backgrounds/darkPurple.png")
	s.Player = s.loadImageFromFS("PNG/playerShip1_orange.png")
	s.Meteors = s.loadMultipleImagesFromFS("PNG/Meteors/*.png")
	s.Bullets = &BulletSprites{
		Basic: &MissileSprites{
			Bullet:    s.loadImageFromFS("PNG/Lasers/laserBlue.png"),
			Explosion: s.loadImageFromFS("PNG/Lasers/laserBlueExplosion.png"),
		},
		Upgrade: &MissileSprites{
			Bullet:    s.loadImageFromFS("PNG/Lasers/Upgrades/laserGreenUpgrade.png"),
			Explosion: s.loadImageFromFS("PNG/Lasers/Upgrades/laserGreenUpgradeExplosion.png"),
		},
		Missile: &MissileSprites{
			Bullet:    s.loadImageFromFS("PNG/Lasers/Upgrades/laserRedMissile.png"),
			Explosion: s.loadImageFromFS("PNG/Lasers/Upgrades/laserRedMissileExplosion.png"),
		},
	}

	return s
}

package assets

import (
	"embed"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed Backgrounds/* PNG/* Fonts/*
var assetsFS embed.FS

type Sprites struct {
	Background *ebiten.Image
	Player     *ebiten.Image
	Meteors    []*ebiten.Image
	Bullets    *BulletSprites
}

type GameFonts struct {
	ScoreFont font.Face
}

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

func loadFont(fontName string, Size float64) font.Face {
	fontFile, err := assetsFS.ReadFile(fontName)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(fontFile)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    Size,
		DPI:     Size * ebiten.DeviceScaleFactor(),
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
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

func LoadFonts() *GameFonts {
	return &GameFonts{
		ScoreFont: loadFont("Fonts/kenvector_future.ttf", 32),
	}
}

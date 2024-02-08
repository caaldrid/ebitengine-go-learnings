package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Background struct {
	sprite       *ebiten.Image
	ScreenWidth  int
	ScreenHeight int
}

func (b *Background) Update() error {
	return nil
}

func (b *Background) Draw(screen *ebiten.Image) {

	xScale := float64(b.ScreenWidth) / float64(b.sprite.Bounds().Size().X)
	yScale := float64(b.ScreenHeight) / float64(b.sprite.Bounds().Size().Y)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(xScale, yScale)
	screen.DrawImage(b.sprite, op)
}

func NewBackground(sprite *ebiten.Image) *Background {
	return &Background{
		sprite:       sprite,
		ScreenWidth:  ScreenWidth,
		ScreenHeight: ScreenHeight,
	}
}

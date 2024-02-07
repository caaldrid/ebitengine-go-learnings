package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func NewBackground(sprite *ebiten.Image) *Background {
	return &Background{
		sprite:       sprite,
		ScreenWidth:  ScreenWidth,
		ScreenHeight: ScreenHeight,
	}
}

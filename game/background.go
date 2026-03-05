package game

import (
	"github.com/caaldrid/ebitengine-go-learnings/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Background struct {
	asset        *assets.Asset
	ScreenWidth  int
	ScreenHeight int
}

func (b *Background) Update() error {
	return nil
}

func (b *Background) Draw(screen *ebiten.Image) {

	xScale := float64(b.ScreenWidth) / float64(b.asset.Sprite.Bounds().Size().X)
	yScale := float64(b.ScreenHeight) / float64(b.asset.Sprite.Bounds().Size().Y)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(xScale, yScale)
	screen.DrawImage(b.asset.Sprite, op)
}

func NewBackground(asset *assets.Asset) *Background {
	return &Background{
		asset:        asset,
		ScreenWidth:  ScreenWidth,
		ScreenHeight: ScreenHeight,
	}
}

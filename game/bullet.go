package game

import (
	"math"

	"github.com/caaldrid/ebitengine-go-learnings/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	asset    *assets.Asset
	angle    float64
	velocity float64
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	if b.asset != nil {
		halfW, halfH := b.asset.CalcCenter()

		op := &ebiten.DrawImageOptions{}

		// Maintain the player sprite on the center by translating in opposing directions before and after rotation
		op.GeoM.Translate(-halfW, -halfH)
		op.GeoM.Rotate(b.angle)
		op.GeoM.Translate(halfW, halfH)

		op.GeoM.Translate(b.asset.Pos.X, b.asset.Pos.Y)
		op.Filter = ebiten.FilterLinear
		screen.DrawImage(b.asset.Sprite, op)
	}
}

func (b *Bullet) Update() error {
	b.asset.Pos.X += math.Sin(b.angle) * b.velocity
	b.asset.Pos.Y += math.Cos(b.angle) * -b.velocity

	return nil
}

func NewBullet(sprite *ebiten.Image, bulletVelocity float64, player *Player) *Bullet {
	asset := &assets.Asset{
		Sprite: sprite,
	}

	halfW, halfH := asset.CalcCenter()
	pHalfW, pHalfH := player.asset.CalcCenter()

	asset.Pos = assets.Vector{
		X: player.asset.Pos.X + pHalfW - halfW,
		Y: player.asset.Pos.Y + pHalfH - halfH,
	}

	return &Bullet{
		asset:    asset,
		angle:    player.angle,
		velocity: bulletVelocity,
	}
}

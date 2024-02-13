package game

import (
	"math"

	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	asset    *assets.Asset
	pos      Vector
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

		op.GeoM.Translate(b.pos.X, b.pos.Y)
		screen.DrawImage(b.asset.Sprite, op)
	}
}

func (b *Bullet) Update() error {
	b.pos.X += math.Sin(b.angle) * b.velocity
	b.pos.Y += math.Cos(b.angle) * -b.velocity

	return nil
}

func NewBullet(asset *assets.Asset, bulletVelocity float64, player *Player) *Bullet {

	halfW, halfH := asset.CalcCenter()
	pHalfW, pHalfH := player.asset.CalcCenter()

	return &Bullet{
		asset: asset,
		angle: player.angle,
		pos: Vector{
			X: player.position.X + pHalfW - halfW,
			Y: player.position.Y + pHalfH - halfH,
		},
		velocity: bulletVelocity,
	}
}

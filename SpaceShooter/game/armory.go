package game

import (
	"math"
	"time"

	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	bulletSpeedPerSecond = 350.0
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

func (b *Bullet) Update(maxXBound, maxYBound float64) error {
	canMoveX := b.pos.X > 0 && b.pos.X < maxXBound
	canMoveY := b.pos.Y > 0 && b.pos.Y < maxYBound

	if canMoveX && canMoveY {
		b.pos.X += math.Sin(b.angle) * b.velocity
		b.pos.Y += math.Cos(b.angle) * -b.velocity
	} else {
		b.asset = nil
	}
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

type Armory struct {
	bulletAssets *assets.BulletAssets
	bullets      []*Bullet
	bulletTimer  *Timer
	hasUpgrade   bool
	maxXBound    float64
	maxYBound    float64
}

func (a *Armory) Update(player *Player) error {
	a.bulletTimer.Update()

	// Check if we can spawn a bullet
	if a.bulletTimer.Completed() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		a.bulletTimer.Reset()

		asset := a.bulletAssets.Basic
		velocity := bulletSpeedPerSecond / float64(ebiten.TPS())
		if a.hasUpgrade {
			asset = a.bulletAssets.Upgrade
			velocity *= 2
		}

		a.bullets = append(a.bullets, NewBullet(asset, velocity, player))
	}

	for _, bullet := range a.bullets {
		err := bullet.Update(a.maxXBound, a.maxYBound)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Armory) Draw(screen *ebiten.Image) {
	for _, bullet := range a.bullets {
		bullet.Draw(screen)
	}
}

func NewArmory(assets *assets.BulletAssets, maxXBound, maxYBound float64) *Armory {
	return &Armory{
		bulletAssets: assets,
		bullets:      make([]*Bullet, 0),
		bulletTimer:  NewTimer(500 * time.Millisecond),
		hasUpgrade:   false,
		maxXBound:    maxXBound,
		maxYBound:    maxYBound,
	}
}

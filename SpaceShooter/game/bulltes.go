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
	bulletAssets *assets.BulletAssets
	asset        *assets.Asset
	timer        *Timer
	pos          Vector
	angle        float64
	isUpgraded   bool
	velocity     float64
	maxXBound    int
	maxYBound    int
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

func (b *Bullet) Update(player *Player) error {
	b.timer.Update()
	basicVelocity := bulletSpeedPerSecond / float64(ebiten.TPS())

	// Check if we can spawn a bullet
	if b.timer.Completed() && b.asset == nil && ebiten.IsKeyPressed(ebiten.KeySpace) {
		b.timer.Reset()

		if b.isUpgraded {
			b.asset = b.bulletAssets.Upgrade
			b.velocity = basicVelocity * 2 // Double the velocity
		} else {
			b.asset = b.bulletAssets.Basic
			b.velocity = basicVelocity
		}

		b.angle = player.angle

		halfW, halfH := b.asset.CalcCenter()
		pHalfW, pHalfH := player.asset.CalcCenter()

		b.pos.X = player.position.X + pHalfW - halfW
		b.pos.Y = player.position.Y + pHalfH - halfH

	} else if b.asset != nil { // Move the bullet if the asset has be assigned
		canMoveX := b.pos.X > 0 && b.pos.X < float64(b.maxXBound)
		canMoveY := b.pos.Y > 0 && b.pos.Y < float64(b.maxYBound)

		if canMoveX && canMoveY {
			b.pos.X += math.Sin(b.angle) * b.velocity
			b.pos.Y += math.Cos(b.angle) * -b.velocity
		} else {
			b.asset = nil
		}
	}
	return nil
}

func NewBullet(bulletAssets *assets.BulletAssets, ScreenWidth, ScreenHeight int) *Bullet {
	return &Bullet{
		bulletAssets: bulletAssets,
		asset:        nil,
		timer:        NewTimer(500 * time.Millisecond),
		isUpgraded:   false,
		maxXBound:    ScreenWidth,
		maxYBound:    ScreenHeight,
	}
}

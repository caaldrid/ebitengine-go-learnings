package game

import (
	"time"

	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	bulletAssets *assets.BulletAssets
	asset        *assets.Asset
	timer        *Timer
	position     Vector
	angle        float64
	isUpgraded   bool
	hasMissile   bool
	speed        float64
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

		op.GeoM.Translate(b.position.X, b.position.Y)
		screen.DrawImage(b.asset.Sprite, op)
	}
}

func (b *Bullet) Update(player *Player) error {
	b.timer.Update()
	// canMoveX := b.position.X > 0 && b.position.X < float64(b.maxXBound)
	// canMoveY := b.position.Y > 0 && b.position.Y < float64(b.maxYBound)

	// Check if we can spawn a bullet
	if b.timer.Completed() && b.asset == nil && ebiten.IsKeyPressed(ebiten.KeySpace) {
		b.timer.Reset()

		if b.isUpgraded {
			b.asset = b.bulletAssets.Upgrade
		} else {
			b.asset = b.bulletAssets.Basic
		}

		b.angle = player.angle

		halfW, halfH := b.asset.CalcCenter()
		pHalfW, pHalfH := player.asset.CalcCenter()

		b.position.X = player.position.X + pHalfW - halfW
		b.position.Y = player.position.Y + pHalfH - halfH

	}
	return nil
}

func NewBullet(bulletAssets *assets.BulletAssets, ScreenWidth, ScreenHeight int) *Bullet {
	return &Bullet{
		bulletAssets: bulletAssets,
		asset:        nil,
		timer:        NewTimer(500 * time.Millisecond),
		isUpgraded:   false,
		hasMissile:   false,
		maxXBound:    ScreenWidth,
		maxYBound:    ScreenHeight,
	}
}

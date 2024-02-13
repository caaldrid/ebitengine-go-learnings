package game

import (
	"time"

	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	bulletSpeedPerSecond = 350.0
)

type Armory struct {
	bulletSprites *assets.BulletSprites
	bullets       []*Bullet
	bulletTimer   *Timer
	hasUpgrade    bool
	maxXBound     float64
	maxYBound     float64
}

func (a *Armory) Update(player *Player) error {
	a.bulletTimer.Update()

	// Check if we can spawn a bullet
	if a.bulletTimer.Completed() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		a.bulletTimer.Reset()

		asset := a.bulletSprites.Basic
		velocity := bulletSpeedPerSecond / float64(ebiten.TPS())
		if a.hasUpgrade {
			asset = a.bulletSprites.Upgrade
			velocity *= 2
		}

		a.bullets = append(a.bullets, NewBullet(asset, velocity, player))
	}

	if len(a.bullets) > 0 {
		inBoundBullets := make([]*Bullet, 0)
		for _, bullet := range a.bullets {
			canMoveX := bullet.asset.Pos.X > 0 && bullet.asset.Pos.X < a.maxXBound
			canMoveY := bullet.asset.Pos.Y > 0 && bullet.asset.Pos.Y < a.maxYBound
			if canMoveX && canMoveY {
				err := bullet.Update()
				if err != nil {
					return err
				}

				// Populate new slice with the bullets that could still be moved
				// This will allow us to not keep pointers to bullets that out of bounds
				inBoundBullets = append(inBoundBullets, bullet)
			}
		}

		a.bullets = inBoundBullets
	}

	return nil
}

func (a *Armory) Draw(screen *ebiten.Image) {
	for _, bullet := range a.bullets {
		bullet.Draw(screen)
	}
}

func NewArmory(sprites *assets.BulletSprites, maxXBound, maxYBound float64) *Armory {
	return &Armory{
		bulletSprites: sprites,
		bullets:       make([]*Bullet, 0),
		bulletTimer:   NewTimer(500 * time.Millisecond),
		hasUpgrade:    false,
		maxXBound:     maxXBound,
		maxYBound:     maxYBound,
	}
}

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

	if len(a.bullets) > 0 {
		inBoundBullets := make([]*Bullet, 0)
		for _, bullet := range a.bullets {
			canMoveX := bullet.pos.X > 0 && bullet.pos.X < a.maxXBound
			canMoveY := bullet.pos.Y > 0 && bullet.pos.Y < a.maxYBound
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

package game

import (
	"math"
	"math/rand"
	"time"

	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	bulletSpeedPerSecond = 350.0
)

type explosion struct {
	asset    *assets.Asset
	timer    *Timer
	rotation float64
	scale    float64
}

type Armory struct {
	bulletSprites *assets.BulletSprites
	bullets       []*Bullet
	explosions    []*explosion
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

		asset := a.bulletSprites.Basic.Bullet
		velocity := bulletSpeedPerSecond / float64(ebiten.TPS())
		if a.hasUpgrade {
			asset = a.bulletSprites.Upgrade.Bullet
			velocity *= 2
		}

		a.bullets = append(a.bullets, NewBullet(asset, velocity, player))
	}

	if len(a.bullets) > 0 {
		for i, bullet := range a.bullets {
			canMoveX := bullet.asset.Pos.X > 0 && bullet.asset.Pos.X < a.maxXBound
			canMoveY := bullet.asset.Pos.Y > 0 && bullet.asset.Pos.Y < a.maxYBound
			if canMoveX && canMoveY {
				err := bullet.Update()
				if err != nil {
					return err
				}
			} else {
				// This will allow us to not keep pointers to bullets that out of bounds
				a.bullets = append(a.bullets[:i], a.bullets[i+1:]...)
			}
		}
	}

	// Update Explosions
	for i, explosion := range a.explosions {
		explosion.timer.Update()
		if explosion.timer.Completed() {
			// Delete the explosion from the system
			a.explosions = append(a.explosions[:i], a.explosions[i+1:]...)
		}

		// Rotate any given explosions still around
		if len(a.explosions) > 0 {
			velocity := 0.25 + rand.Float64()*1.5
			// calculate the spin of the meteor
			explosion.rotation += math.Pi / float64(ebiten.TPS()) * velocity
		}
	}

	return nil
}

func (a *Armory) Draw(screen *ebiten.Image) {
	// Draw bullets
	for _, bullet := range a.bullets {
		bullet.Draw(screen)
	}

	// Draw explosions
	for _, explosion := range a.explosions {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(explosion.scale, explosion.scale)
		op.GeoM.Translate(explosion.asset.Pos.X, explosion.asset.Pos.Y)
		screen.DrawImage(explosion.asset.Sprite, op)
	}
}

func (a *Armory) HandleTargetHit(bulletIndex int, meteorHit *assets.Asset) {

	// Delete the bullet from the system
	a.bullets = append(a.bullets[:bulletIndex], a.bullets[bulletIndex+1:]...)

	// Create a new explosion
	explosionSprite := a.bulletSprites.Basic.Explosion
	if a.hasUpgrade {
		explosionSprite = a.bulletSprites.Upgrade.Explosion
	}

	// Calculate relative scale of explosion from meteor sizeff
	meteorScale :=

		float64(meteorHit.Sprite.Bounds().Dx()) / float64(explosionSprite.Bounds().Dx())
	a.explosions = append(a.explosions, &explosion{
		asset: &assets.Asset{
			Sprite: explosionSprite,
			Pos:    meteorHit.Pos,
		},
		rotation: 0,
		scale:    meteorScale,
		timer:    NewTimer(250 * time.Millisecond),
	})
}

func NewArmory(sprites *assets.BulletSprites, maxXBound, maxYBound float64) *Armory {
	return &Armory{
		bulletSprites: sprites,
		bullets:       make([]*Bullet, 0),
		explosions:    make([]*explosion, 0),
		bulletTimer:   NewTimer(500 * time.Millisecond),
		hasUpgrade:    false,
		maxXBound:     maxXBound,
		maxYBound:     maxYBound,
	}
}

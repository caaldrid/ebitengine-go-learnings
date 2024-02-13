package game

import (
	"math"
	"math/rand"
	"time"

	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

var target assets.Vector

type meteor struct {
	asset    *assets.Asset
	rotation float64
}

func (m *meteor) move() {
	// Randomized velocity
	velocity := 0.25 + rand.Float64()*1.5
	halfW, halfH := m.asset.CalcCenter()

	// Direction is the target minus the current position
	direction := assets.Vector{
		X: target.X - (m.asset.Pos.X + halfW),
		Y: target.Y - (m.asset.Pos.Y + halfH),
	}

	// Normalize the vector — get just the direction without the length
	normalizedDirection := direction.Normalize()

	// Update the meteor's position
	m.asset.Pos.X += normalizedDirection.X * velocity
	m.asset.Pos.Y += normalizedDirection.Y * velocity

	// calculate the spin of the meteor
	m.rotation += math.Pi / float64(ebiten.TPS()) * velocity
}

type Meteors struct {
	meteorsSprites []*ebiten.Image
	meteors        []*meteor
	timer          *Timer
	spawnRadious   float64
	player         *Player
}

func (me *Meteors) Update() error {
	me.timer.Update()

	// Spawn a new meteor around the player if the timer has completed
	if me.timer.Completed() {
		me.timer.Reset()

		newMetorSprite := me.meteorsSprites[rand.Intn(len(me.meteorsSprites))]

		// Calcuate where in the circle the meteor will spawn
		angle := rand.Float64() * 2 * math.Pi
		posX := target.X + math.Cos(angle)*me.spawnRadious
		posY := target.Y + math.Sin(angle)*me.spawnRadious

		newMetor := &meteor{
			asset: &assets.Asset{
				Sprite: newMetorSprite,
				Pos: assets.Vector{
					X: posX,
					Y: posY,
				},
			},
		}

		me.meteors = append(me.meteors, newMetor)

	}

	// Move meteors
	for _, meteor := range me.meteors {
		meteor.move()
	}
	return nil
}

func (me *Meteors) Draw(screen *ebiten.Image) {
	for _, metor := range me.meteors {
		op := &ebiten.DrawImageOptions{}

		// Rotate
		halfW, halfH := metor.asset.CalcCenter()
		op.GeoM.Translate(-halfW, -halfH)
		op.GeoM.Rotate(metor.rotation)
		op.GeoM.Translate(halfW, halfH)

		op.GeoM.Translate(metor.asset.Pos.X, metor.asset.Pos.Y)
		screen.DrawImage(metor.asset.Sprite, op)
	}
}

func NewMetors(sprites []*ebiten.Image, player *Player, ScreenWidth int) *Meteors {
	playerHalfW, playerHalfH := player.asset.CalcCenter()
	target.X = player.asset.Pos.X + playerHalfW
	target.Y = player.asset.Pos.Y + playerHalfH

	return &Meteors{
		meteorsSprites: sprites,
		timer:          NewTimer(3000 * time.Millisecond),
		spawnRadious:   float64(ScreenWidth) / 2.0,
		meteors:        make([]*meteor, 0),
		player:         player,
	}
}

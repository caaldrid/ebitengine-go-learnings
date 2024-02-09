package game

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type meteor struct {
	sprite *ebiten.Image
	pos    Vector
}

type Meteors struct {
	sprites      []*ebiten.Image
	meteors      []*meteor
	timer        *Timer
	spawnRadious float64
	player       *Player
}

func (me *Meteors) Update() error {
	me.timer.Update()
	if me.timer.Completed() {
		me.timer.Reset()

		// Spawn a new meteor around the player
		newMetorSprite := me.sprites[rand.Intn(len(me.sprites))]

		angle := rand.Float64() * 2 * math.Pi
		newMetor := &meteor{
			sprite: newMetorSprite,
			pos: Vector{
				X: me.player.position.X + math.Cos(angle)*me.spawnRadious,
				Y: me.player.position.Y + math.Sin(angle)*me.spawnRadious,
			},
		}

		me.meteors = append(me.meteors, newMetor)
	}
	return nil
}

func (me *Meteors) Draw(screen *ebiten.Image) {
	for _, metor := range me.meteors {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(metor.pos.X, metor.pos.Y)
		screen.DrawImage(metor.sprite, op)
	}
}

func NewMetors(sprites []*ebiten.Image, player *Player, ScreenWidth int) *Meteors {
	return &Meteors{
		sprites:      sprites,
		timer:        NewTimer(3 * time.Second),
		spawnRadious: float64(ScreenWidth) / 2.0,
		meteors:      make([]*meteor, 0),
		player:       player,
	}
}

package game

import (
	"math"

	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	sprite   *ebiten.Image
	position Vector
	rotation float64
}

func (p *Player) Update() error {
	//  rotate 180° per second.
	speed := math.Pi / float64(ebiten.TPS())

	// Go left
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.rotation -= speed
	}

	// Go right
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		p.rotation += speed
	}

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	halfW, halfH := assets.CalcCenter(p.sprite)

	op := &ebiten.DrawImageOptions{}

	// Maintain the player sprite on the center by translating in opposing directions before and after rotation
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

func NewPlayer(sprite *ebiten.Image, ScreenWidth, ScreenHeight int) *Player {
	halfW, halfH := assets.CalcCenter(sprite)
	// Set Player position to the center of the screen offseted by the plater center
	pos := Vector{
		X: float64(ScreenWidth)/2 - halfW,
		Y: float64(ScreenHeight)/2 - halfH,
	}

	return &Player{
		sprite:   sprite,
		position: pos,
	}
}

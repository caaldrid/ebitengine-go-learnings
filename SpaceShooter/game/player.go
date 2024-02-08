package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X float64
	Y float64
}

type Player struct {
	sprite   *ebiten.Image
	position Vector
	rotation float64
}

func calcCenter(sprite *ebiten.Image) (float64, float64) {
	// Find the center of the sprite
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	return halfW, halfH
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
	halfW, halfH := calcCenter(p.sprite)

	op := &ebiten.DrawImageOptions{}

	// Maintain the player sprite on the center by translating in opposing directions before and after rotation
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

func NewPlayer(sprite *ebiten.Image, ScreenWidth, ScreenHeight int) *Player {
	halfW, halfH := calcCenter(sprite)
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

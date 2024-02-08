package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X float64
	Y float64
}

type Player struct {
	sprite   *ebiten.Image
	position Vector
}

func (p *Player) Update() error {

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

func NewPlayer(sprite *ebiten.Image, ScreenWidth, ScreenHeight int) *Player {
	// Find the center of the sprite
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

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

package assets

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Normalize() Vector {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return Vector{v.X / magnitude, v.Y / magnitude}
}

type Asset struct {
	Sprite *ebiten.Image
	Pos    Vector
}

func (a *Asset) CalcCenter() (float64, float64) {
	// Find the center of the sprite
	bounds := a.Sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	return halfW, halfH
}

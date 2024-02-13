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

func (a *Asset) MaxX() float64 {
	return a.Pos.X + float64(a.Sprite.Bounds().Dx())
}

func (a *Asset) MaxY() float64 {
	return a.Pos.Y + float64(a.Sprite.Bounds().Dy())
}

func (a *Asset) Intersects(other *Asset) bool {
	return a.Pos.X <= other.MaxX() &&
		other.Pos.X <= a.MaxX() &&
		a.Pos.Y <= other.MaxY() &&
		other.Pos.Y <= a.MaxY()
}

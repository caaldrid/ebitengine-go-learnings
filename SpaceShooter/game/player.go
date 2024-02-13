package game

import (
	"math"

	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	asset    *assets.Asset
	Armory   *Armory
	position Vector
	angle    float64
}

func (p *Player) Update() error {
	//  rotate 180° per second.
	speed := math.Pi / float64(ebiten.TPS())

	// Go left
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.angle -= speed
	}

	// Go right
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		p.angle += speed
	}

	return p.Armory.Update(p)

}

func (p *Player) Draw(screen *ebiten.Image) {
	// Draw bullets first so that its under the player sprite
	p.Armory.Draw(screen)

	halfW, halfH := p.asset.CalcCenter()

	op := &ebiten.DrawImageOptions{}

	// Maintain the player sprite on the center by translating in opposing directions before and after rotation
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.angle)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.asset.Sprite, op)

}

func NewPlayer(playerAsset *assets.Asset, bulletAssets *assets.BulletAssets, ScreenWidth, ScreenHeight int) *Player {
	halfW, halfH := playerAsset.CalcCenter()
	// Set Player position to the center of the screen offseted by the plater center
	pos := Vector{
		X: float64(ScreenWidth)/2 - halfW,
		Y: float64(ScreenHeight)/2 - halfH,
	}

	return &Player{
		asset:    playerAsset,
		Armory:   NewArmory(bulletAssets, float64(ScreenWidth), float64(ScreenHeight)),
		position: pos,
		angle:    0,
	}
}

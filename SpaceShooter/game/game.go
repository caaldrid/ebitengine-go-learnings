package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Background *Background
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Background.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Background.ScreenWidth, g.Background.ScreenHeight
}

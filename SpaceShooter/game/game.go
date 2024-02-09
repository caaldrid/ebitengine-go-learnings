package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Background *Background
	Player     *Player
	Meteors    *Meteors
}

func (g *Game) Update() error {
	err := g.Player.Update()
	if err != nil {
		return err
	}

	err = g.Meteors.Update()
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Background.Draw(screen)
	g.Player.Draw(screen)
	g.Meteors.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Background.ScreenWidth, g.Background.ScreenHeight
}

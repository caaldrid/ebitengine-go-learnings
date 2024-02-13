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

	for i, meteor := range g.Meteors.meteors {

		// Handle the collision between meteors and bullets
		for j, bullet := range g.Player.Armory.bullets {
			if meteor.asset.Intersects(bullet.asset) {
				g.Meteors.meteors = append(g.Meteors.meteors[:i], g.Meteors.meteors[i+1:]...)
				g.Player.Armory.bullets = append(g.Player.Armory.bullets[:j], g.Player.Armory.bullets[j+1:]...)
			}
		}
	}

	// Need to iterate through any meteors left after taking into account collisions with bullets
	for i, meteor := range g.Meteors.meteors {
		// Handle the collision between meteors and player
		if meteor.asset.Intersects(g.Player.asset) {
			g.Meteors.meteors = append(g.Meteors.meteors[:i], g.Meteors.meteors[i+1:]...) // Delete the Meteor
		}
	}
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

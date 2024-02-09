package main

import (
	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/assets"
	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	assets := assets.NewAssets()
	background := game.NewBackground(assets.Background)
	player := game.NewPlayer(assets.Player, background.ScreenWidth, background.ScreenHeight)
	meteors := game.NewMetors(assets.Meteors, player, background.ScreenWidth)

	g := &game.Game{Background: background, Player: player, Meteors: meteors}

	ebiten.SetWindowSize(background.ScreenWidth, background.ScreenHeight)
	ebiten.SetWindowTitle("Space Shooter (Ebitengine Learning)")

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}

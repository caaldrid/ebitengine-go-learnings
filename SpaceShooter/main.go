package main

import (
	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/assets"
	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	assets := assets.NewAssets()
	background := game.NewBackground(assets.Background)

	g := &game.Game{Background: background}

	ebiten.SetWindowSize(background.ScreenWidth, background.ScreenHeight)
	ebiten.SetWindowTitle("Space Shooter (Ebitengine Learning)")

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}

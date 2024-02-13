package main

import (
	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/assets"
	"github.com/caaldrid/ebitengine-go-learnings/SpaceShooter/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	sprites := assets.LoadSprites()

	background := game.NewBackground(&assets.Asset{
		Sprite: sprites.Background,
	})
	player := game.NewPlayer(&assets.Asset{
		Sprite: sprites.Player,
	}, sprites.Bullets, background.ScreenWidth, background.ScreenHeight)
	meteors := game.NewMetors(sprites.Meteors, player, background.ScreenWidth)

	g := &game.Game{Background: background, Player: player, Meteors: meteors}

	ebiten.SetWindowSize(background.ScreenWidth, background.ScreenHeight)
	ebiten.SetWindowTitle("Space Shooter (Ebitengine Learning)")

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}

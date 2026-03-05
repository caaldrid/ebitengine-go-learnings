package main

import (
	"github.com/caaldrid/ebitengine-go-learnings/assets"
	"github.com/caaldrid/ebitengine-go-learnings/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	sprites := assets.LoadSprites()
	g := game.NewGame(sprites)

	ebiten.SetWindowSize(g.Background.ScreenWidth, g.Background.ScreenHeight)
	ebiten.SetWindowTitle("Space Shooter (Ebitengine Learning)")

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}

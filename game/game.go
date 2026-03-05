package game

import (
	"fmt"
	"image/color"
	"math"

	"github.com/caaldrid/ebitengine-go-learnings/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Game struct {
	Background *Background
	Player     *Player
	Meteors    *Meteors
	score      int
	Fonts      *assets.GameFonts
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
				g.Player.Armory.HandleTargetHit(j, meteor.asset)
				g.Meteors.meteors = append(g.Meteors.meteors[:i], g.Meteors.meteors[i+1:]...)
				g.score++
			}
		}
	}

	// Need to iterate through any meteors left after taking into account collisions with bullets
	for _, meteor := range g.Meteors.meteors {
		// Handle the collision between meteors and player
		if meteor.asset.Intersects(g.Player.asset) {
			g.Reset()
			break
		}
	}
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Background.Draw(screen)
	g.Player.Draw(screen)
	g.Meteors.Draw(screen)

	// Draw score on screen
	scoreStr := fmt.Sprintf("\nScore: %06d", g.score)
	gameScoreBounds, _ := font.BoundString(g.Fonts.ScoreFont, scoreStr)
	scoreDrawXPos := g.Background.ScreenWidth/2 - int(gameScoreBounds.Max.X.Floor()/2)
	scoreDrawYPos := int(math.Abs(float64(gameScoreBounds.Min.Y.Floor())))
	text.Draw(screen, scoreStr, g.Fonts.ScoreFont, scoreDrawXPos, scoreDrawYPos, color.White)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Background.ScreenWidth, g.Background.ScreenHeight
}

func (g *Game) Reset() {
	newGameState := NewGame(assets.LoadSprites())

	g.Background = newGameState.Background
	g.Player = newGameState.Player
	g.Meteors = newGameState.Meteors
	g.score = 0
}

func NewGame(sprites *assets.Sprites) *Game {
	background := NewBackground(&assets.Asset{
		Sprite: sprites.Background,
	})
	player := NewPlayer(&assets.Asset{
		Sprite: sprites.Player,
	}, sprites.Bullets, background.ScreenWidth, background.ScreenHeight)
	meteors := NewMetors(sprites.Meteors, player, background.ScreenWidth)

	return &Game{Background: background, Player: player, Meteors: meteors, Fonts: assets.LoadFonts()}
}

package game

import (
	"errors"

	//	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const GameWidth = 1920
const GameHeight = 1080

var (
	regularTermination = errors.New("regular termination")
)

type Game struct {
	count int
	board *Board
}

func NewGame() *Game {
	game := new(Game)
	game.board = NewBoard()
	return game
}

func (g *Game) Update() error {
	g.count++

	g.board.Update()
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {

		return regularTermination
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.board.MovePaddle(0, -5.0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.board.MovePaddle(0, 5.0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyI) {
		g.board.MovePaddle(1, -5.0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		g.board.MovePaddle(1, 5.0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//	fw, fh := ebiten.ScreenSizeInFullscreen()
	//sw, sh := screen.Size()
	g.board.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GameWidth, GameHeight
}

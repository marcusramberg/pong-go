package game

import (
	"errors"

	//	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

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

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

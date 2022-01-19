package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Board struct {
	count   int
	paddles []float64
}

func NewBoard() *Board {
	board := new(Board)
	board.paddles = []float64{10.0, 10.0}
	return board
}

func (b *Board) MovePaddle(paddle int, pos float64) {
	b.paddles[paddle] += pos
}

func (b *Board) Draw(screen *ebiten.Image) {
	sw, sh := screen.Size()
	center := float64(sw) / 2
	scale := ebiten.DeviceScaleFactor()
	for y := float64(30); y < float64(sh)-110; y = y + 110 {
		ebitenutil.DrawRect(screen, center-5.0, y, 10, 70, color.White)
	}
	msgs := []string{
		"esc to quit",
	}
	for i, msg := range msgs {
		text.Draw(screen, msg, mplusFont, int(50*scale), int(100+float64(i)*40*scale), color.White)
	}
	// Just draw two paddles for now
	//for i, paddle := range b.paddles {
	//}
	ebitenutil.DrawRect(screen, 20*scale, b.paddles[0]*scale, 15, 100, color.White)
	ebitenutil.DrawRect(screen, float64(sw)-20*scale, b.paddles[1]*scale, 15, 100, color.White)

}

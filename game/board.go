package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Board struct {
	count int
}

func NewBoard() *Board {
	return new(Board)
}

func (b *Board) Draw(screen *ebiten.Image) {
	sw, sh := screen.Size()
	center := float64(sw)/2
	for y := float64(30); y < float64(sh)-110; y = y+110 {
		ebitenutil.DrawRect(screen, center-5.0 , y, 10 , 70, color.White)
	}
	text.Draw(screen, msg, mplusFont, int(50*scale), int(100+float64(i)*40*scale), color.White)


}

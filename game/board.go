package game

import (
	"image/color"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	smallFont font.Face
	bigFont   font.Face
)

func initFonts() {
	squareFont, err := ioutil.ReadFile("assets/Square.ttf")
	if err != nil {
		log.Fatal(err)
	}
	tt, err := opentype.Parse(squareFont)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	smallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24 * ebiten.DeviceScaleFactor(),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	bigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48 * ebiten.DeviceScaleFactor(),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Board ...
type Board struct {
	count   int
	paddles []float64
	scores  []int
	ball    *Ball
}

// NewBoard
func NewBoard() *Board {
	board := new(Board)
	board.ball = NewBall(false)
	board.paddles = []float64{10.0, 10.0}
	board.scores = []int{0, 0}
	initFonts()
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
	text.Draw(screen, "esc to quit", smallFont, sw-int(400.0*scale), sh-int(30.0*scale), color.RGBA{0xbb, 0xbb, 0xbb, 0xff})
	text.Draw(screen, strconv.Itoa(b.scores[0]), bigFont, int(center)-int(70.0*scale), int(70.0*scale), color.RGBA{0xbb, 0xbb, 0xbb, 0xff})
	text.Draw(screen, strconv.Itoa(b.scores[1]), bigFont, int(center)+int(40.0*scale), int(70.0*scale), color.RGBA{0xbb, 0xbb, 0xbb, 0xff})
	// Just draw two paddles for now
	//for i, paddle := range b.paddles {
	//}
	ebitenutil.DrawRect(screen, 50*scale, b.paddles[0]*scale, 15, 100, color.White)
	ebitenutil.DrawRect(screen, float64(sw)-60*scale, b.paddles[1]*scale, 15, 100, color.White)
	b.ball.draw(screen)
}

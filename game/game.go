package game

import (
	"errors"
	"io/ioutil"
	"log"

	//	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	regularTermination = errors.New("regular termination")
	mplusFont          font.Face
)

func initFont() {
	squareFont, err := ioutil.ReadFile("assets/Square.ttf")
	if err != nil {
		log.Fatal(err)
	}
	tt, err := opentype.Parse(squareFont)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24 * ebiten.DeviceScaleFactor(),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	count int
	board *Board
}

func NewGame() *Game {
	game := new(Game)
	game.board = NewBoard()
	initFont()
	return game
}

func (g *Game) Update() error {
	g.count++
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyEscape):
		return regularTermination
	case ebiten.IsKeyPressed(ebiten.KeyW):
		g.board.MovePaddle(0, -5.0)
	case ebiten.IsKeyPressed(ebiten.KeyS):
		g.board.MovePaddle(0, 5.0)
	case ebiten.IsKeyPressed(ebiten.KeyI):
		g.board.MovePaddle(1, -5.0)
	case ebiten.IsKeyPressed(ebiten.KeyK):
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

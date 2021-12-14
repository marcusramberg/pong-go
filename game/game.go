package game

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
		regularTermination = errors.New("regular termination")
		mplusFont font.Face
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
}

func NewGame() *Game {
	game := new(Game)
	initFont()
	return game
}

func (g *Game) Update() error {
	g.count++
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return regularTermination
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scale := ebiten.DeviceScaleFactor()
	fw, fh := ebiten.ScreenSizeInFullscreen()
	sw, sh := screen.Size()
	msgs := []string{
		"esc to quit",
		fmt.Sprintf("Screen size in fullscreen: %d, %d", fw, fh),
		fmt.Sprintf("Game's screen size: %d, %d", sw, sh),
		fmt.Sprintf("Device scale factor: %0.2f", scale),
	}

	for i, msg := range msgs {
		text.Draw(screen, msg, mplusFont, int(50*scale), int(100+float64(i)*40*scale), color.White)
	}
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

package main

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/marcusramberg/pong-go/game"
)

var regularTermination = errors.New("regular termination")

func main() {

	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(game.NewGame()); err != nil && err != regularTermination {
		log.Fatal(err)
	}
}

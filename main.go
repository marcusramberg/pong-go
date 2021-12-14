package main

import (
	"errors"
	"log"

	"github.com/marcusramberg/pong-go/game"

	"github.com/hajimehoshi/ebiten/v2"
)

var regularTermination = errors.New("regular termination")

func main() {

	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(game.NewGame()); err != nil && err != regularTermination {
		log.Fatal(err)
	}
}

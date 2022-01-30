package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const GameWidth = 1920;
const gameHeight = 1080;
type Ball struct {
	x, y      int
	x_velocity float64
	y_velocity float64
	size      int
}

func NewBall(right bool) *Ball {
	x := 0
	x_velocity := rand.Float64()*8.0
	y_velocity := (rand.Float64()*16.0)-8
	if right { x = 1920 }
	y := rand.Intn(1080)
	ball := Ball{x: x, y: y, y_velocity: y_velocity, x_velocity: x_velocity, size: 3}
	return &ball
}

func (b *Ball) update() {
	b.x += int(b.x_velocity)
	b.y += int(b.y_velocity)
	
}

func (b *Ball) draw(screen *ebiten.Image) {
}

func (b *Ball) resetPos(x int, y int, dir float64) {
	b.x = x
	b.y = y
}

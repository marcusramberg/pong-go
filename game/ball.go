package game

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
	x, y       float64
	x_velocity float64
	y_velocity float64
	size       float64
	sfx        *Sfx
}

func NewBall(right bool, sfx *Sfx) *Ball {
	x := 0.0
	x_velocity := rand.Float64()*4.0 + 6.0
	y_velocity := rand.Float64() * 8.0
	if right {
		x = GameWidth - 1.0
		x_velocity = 0 - x_velocity
	}
	y := rand.Float64() * float64(GameHeight)
	ball := Ball{x: x, y: y, y_velocity: y_velocity, x_velocity: x_velocity, size: 8.0, sfx: sfx}
	return &ball
}


func (b *Ball) CheckHit(paddles []float64) {
	if b.x_velocity > 0 {
		if b.x >= GameWidth-65.0 && b.x <= GameWidth-50.0 && b.y >= paddles[1] && b.y <= paddles[1]+200.0 {
			b.x_velocity = 0 - b.x_velocity + 0.3
			factor := (b.y-paddles[1])/100.0+0.5
			b.y_velocity*=factor
			err := b.sfx.Play("plop")
			if err != nil { log.Fatal(err) }
		}
	} else {
		if b.x >= 50.0 && b.x <= 65.0 && b.y >= paddles[0] && b.y <= paddles[0]+200.0 {
			b.x_velocity = 0 - b.x_velocity + 0.8
			factor := (b.y-paddles[0])/100.0+0.5
			b.y_velocity*=factor
			err := b.sfx.Play("plop")
			if err != nil { log.Fatal(err) }
		}
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.x, b.y, b.size, b.size, color.White)
}


func (b *Ball) Update() {
	if b.y < 0.0 || b.y+b.size > GameHeight {
		b.y_velocity = 0.0 - b.y_velocity
		err := b.sfx.Play("plop")
		if err != nil { log.Fatal(err) }
	}
	b.x += b.x_velocity
	b.y += b.y_velocity
}

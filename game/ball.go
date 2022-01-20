package game

type Ball struct {
	x, y      int
	speed     float64
	direction float64
	size      int
}

func NewBall() *Ball {
	ball := Ball{x: 30, y: 30, direction: 70.0, speed: 10.0, size: 3}
	return &ball
}

func (b *Ball) resetPos(x int, y int, dir float64) {
	b.x = x
	b.y = y
	b.direction = dir
}

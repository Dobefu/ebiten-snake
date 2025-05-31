package main

import (
	"image/color"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Fruit struct {
	GameObject

	Position Vector2
	snake    *Snake
}

func (f *Fruit) Update() error {
	return nil
}

func (f *Fruit) Tick() {
}

func (f *Fruit) Draw(screen *ebiten.Image) {
	brightness := uint8((math.Sin(float64(game.frame)/10)/2 + .5) * 100)

	vector.DrawFilledCircle(
		screen,
		f.Position.X+16,
		f.Position.Y+16,
		16,
		color.RGBA{
			R: 255,
			G: brightness,
			B: brightness,
		},
		true,
	)
}

func (f *Fruit) RandomizePosition() {
	canPlace := true

	f.Position.X = -99
	f.Position.Y = -99

	go func() {
		var tx, ty float32

		for {
			canPlace = true
			tx = float32(rand.N(20) * 32)
			ty = float32(rand.N(20) * 32)

			if tx == f.snake.Position.X && ty == f.snake.drawPosition.Y {
				canPlace = false
			} else {
				for _, segment := range f.snake.segments {
					if tx == segment.position.X && ty == segment.position.Y {
						canPlace = false
					}
				}
			}

			if canPlace {
				break
			}
		}

		f.Position.X = tx
		f.Position.Y = ty
	}()
}

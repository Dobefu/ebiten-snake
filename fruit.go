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
	f.Position.X = float32(rand.N(21) * 32)
	f.Position.Y = float32(rand.N(21) * 32)
}

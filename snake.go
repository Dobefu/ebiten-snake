package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Snake struct {
	GameObject

	X float32
	Y float32
}

func (s *Snake) Update() error {
	return nil
}

func (s *Snake) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, s.X, s.Y, 32, 32, color.White, true)
}

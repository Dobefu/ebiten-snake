package main

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Snake struct {
	GameObject

	Position Vector2

	length   int
	facing   Direction
	fruit    *Fruit
	segments []Vector2
}

func (s *Snake) Update() error {
	if inpututil.KeyPressDuration(ebiten.KeyH) > 0 &&
		s.facing != Direction(DirectionRight) {
		s.facing = Direction(DirectionLeft)
	}

	if inpututil.KeyPressDuration(ebiten.KeyL) > 0 &&
		s.facing != Direction(DirectionLeft) {
		s.facing = Direction(DirectionRight)
	}

	if inpututil.KeyPressDuration(ebiten.KeyJ) > 0 &&
		s.facing != Direction(DirectionUp) {
		s.facing = Direction(DirectionDown)
	}

	if inpututil.KeyPressDuration(ebiten.KeyK) > 0 &&
		s.facing != Direction(DirectionDown) {
		s.facing = Direction(DirectionUp)
	}

	return nil
}

func (s *Snake) Tick() {
	s.segments = append(s.segments, s.Position)

	if len(s.segments) > s.length {
		s.segments = s.segments[1:]
	}

	s.Position.X += s.facing.X * 32
	s.Position.Y += s.facing.Y * 32

	for _, segment := range s.segments {
		if s.Position.X == segment.X && s.Position.Y == segment.Y {
			os.Exit(0)
		}
	}

	if s.Position.X >= 640 {
		s.Position.X = 0
	}

	if s.Position.X < 0 {
		s.Position.X = 608
	}

	if s.Position.Y >= 640 {
		s.Position.Y = 0
	}

	if s.Position.Y < 0 {
		s.Position.Y = 608
	}

	if s.Position.X == s.fruit.Position.X && s.Position.Y == s.fruit.Position.Y {
		s.fruit.RandomizePosition()
		s.length += 1
	}
}

func (s *Snake) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		s.Position.X,
		s.Position.Y,
		32,
		32,
		color.Gray16{Y: 0xffff},
		true,
	)

	for id, segment := range s.segments {
		vector.DrawFilledRect(
			screen,
			segment.X,
			segment.Y,
			32,
			32,
			color.Gray{Y: 128 + uint8((float32(id)/float32(s.length))*127)},
			true,
		)
	}
}

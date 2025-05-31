package main

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SnakeSegment struct {
	position     Vector2
	drawPosition Vector2
}

type Snake struct {
	GameObject

	Position     Vector2
	drawPosition Vector2

	length   int
	facing   Direction
	fruit    *Fruit
	segments []SnakeSegment
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
	s.segments = append(s.segments, SnakeSegment{
		position:     s.Position,
		drawPosition: s.Position,
	})

	if len(s.segments) > s.length {
		s.segments = s.segments[1:]
	}

	s.Position.X += s.facing.X * 32
	s.Position.Y += s.facing.Y * 32

	for i := range s.segments {
		if i <= 0 {
			continue
		}
	}

	for _, segment := range s.segments {
		if s.Position.X == segment.position.X &&
			s.Position.Y == segment.position.Y {
			os.Exit(0)
		}
	}

	if s.Position.X >= 640 {
		s.Position.X = 0
		s.drawPosition.X = 0
	}

	if s.Position.X < 0 {
		s.Position.X = 608
		s.drawPosition.X = 608
	}

	if s.Position.Y >= 640 {
		s.Position.Y = 0
		s.drawPosition.Y = 0
	}

	if s.Position.Y < 0 {
		s.Position.Y = 608
		s.drawPosition.Y = 608
	}

	if s.Position.X == s.fruit.Position.X && s.Position.Y == s.fruit.Position.Y {
		s.fruit.RandomizePosition()
		s.length += 1
	}
}

func (s *Snake) Draw(screen *ebiten.Image) {
	s.drawPosition.X += (s.Position.X - s.drawPosition.X) / 10
	s.drawPosition.Y += (s.Position.Y - s.drawPosition.Y) / 10

	vector.DrawFilledRect(
		screen,
		s.drawPosition.X,
		s.drawPosition.Y,
		32,
		32,
		color.Gray{Y: 255},
		false,
	)

	for i := range s.segments {
		s.segments[i].drawPosition.X += (s.segments[i].position.X - s.segments[i].drawPosition.X) / 10
		s.segments[i].drawPosition.Y += (s.segments[i].position.Y - s.segments[i].drawPosition.Y) / 10

		vector.DrawFilledRect(
			screen,
			s.segments[i].drawPosition.X,
			s.segments[i].drawPosition.Y,
			32,
			32,
			color.Gray{Y: 128 + uint8((float32(i)/float32(s.length))*127)},
			false,
		)
	}
}

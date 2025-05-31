package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	game  = &Game{tickRate: 30}
	snake = &Snake{
		Position:     Vector2{X: 288, Y: 288},
		drawPosition: Vector2{X: 288, Y: 288},
	}
	fruit = &Fruit{}

	textFace *text.GoTextFace
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.ArcadeN_ttf))

	if err != nil {
		log.Fatal(err)
	}

	textFace = &text.GoTextFace{
		Source: s,
		Size:   16,
	}
}

type Game struct {
	ebiten.Game

	frame       int64
	tickRate    int
	tickIndex   float64
	gameObjects []GameObject
}

func (g *Game) AddGameObject(gameObject GameObject) {
	g.gameObjects = append(g.gameObjects, gameObject)
}

func (g *Game) Update() (err error) {
	g.frame += 1
	g.tickIndex -= ebiten.ActualFPS() / 60

	if g.tickIndex <= 0 {
		g.tickIndex = float64(g.tickRate)

		for _, gameObject := range g.gameObjects {
			gameObject.Tick()
		}
	}

	for _, gameObject := range g.gameObjects {
		err = gameObject.Update()

		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, gameObject := range g.gameObjects {
		gameObject.Draw(screen)
	}

	// UI.
	vector.DrawFilledRect(screen, 0, 0, 640, 64, color.Gray{Y: 30}, false)

	options := &text.DrawOptions{}
	options.GeoM.Translate(16, 24)
	text.Draw(screen, fmt.Sprintf("Score: %d", snake.length), textFace, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 704
}

func main() {
	ebiten.SetWindowSize(640, 704)
	ebiten.SetWindowTitle("Snake")

	snake.fruit = fruit
	fruit.snake = snake
	fruit.RandomizePosition()

	game.AddGameObject(snake)
	game.AddGameObject(fruit)

	err := ebiten.RunGame(game)

	if err != nil {
		log.Fatal(err)
	}
}

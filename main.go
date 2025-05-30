package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	game = &Game{tickRate: 30}
)

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
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func main() {
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Snake")

	snake := &Snake{Position: Vector2{X: 288, Y: 288}}
	fruit := &Fruit{}

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

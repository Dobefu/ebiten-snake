package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ebiten.Game

	gameObjects []GameObject
}

func (g *Game) Update() (err error) {
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
	return 240, 240
}

func main() {
	ebiten.SetWindowSize(480, 480)
	ebiten.SetWindowTitle("Snake")

	game := &Game{}
	game.gameObjects = append(game.gameObjects, &Snake{X: 10, Y: 10})

	err := ebiten.RunGame(game)

	if err != nil {
		log.Fatal(err)
	}
}

package main

import "github.com/hajimehoshi/ebiten/v2"

type GameObject interface {
	Update() (err error)
	Draw(screen *ebiten.Image)
}

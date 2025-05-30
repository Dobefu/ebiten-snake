package main

type Direction Vector2

var (
	DirectionLeft  = Vector2{X: -1, Y: 0}
	DirectionRight = Vector2{X: 1, Y: 0}
	DirectionUp    = Vector2{X: 0, Y: -1}
	DirectionDown  = Vector2{X: 0, Y: 1}
)

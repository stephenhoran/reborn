package game

import (
	"github.com/hajimehoshi/ebiten"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Right:
		return "Right"
	case Down:
		return "Down"
	case Left:
		return "Left"
	default:
		panic("direction should not be reached")
	}
}

type MouseState int

const (
	mouseStateButtonUp MouseState = iota
	mouseStateButtonDown
)

type Input struct {
	mousePosX  int
	mousePosY  int
	mouseState MouseState
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Update() {
	x, y := ebiten.CursorPosition()
	i.mousePosX = x
	i.mousePosY = y
}

func (i *Input) Direction() (Direction, bool) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		return Up, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		return Right, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		return Down, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		return Left, true
	}

	return 0, false
}

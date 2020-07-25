package input

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

func (d Direction) Move() (x, y int) {
	switch d {
	case Up:
		return 0, 1
	case Right:
		return 1, 0
	case Down:
		return 0, -1
	case Left:
		return -1, 0
	}

	return 0, 0
}

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

func (i *Input) MouseX() int {
	return i.mousePosX
}

func (i *Input) MouseY() int {
	return i.mousePosY
}

func (i *Input) MouseLocation() (x, y int) {
	x = i.mousePosX
	y = i.mousePosY

	return
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

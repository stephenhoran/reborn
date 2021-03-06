package input

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Direction int

const (
	Up Direction = iota
	UpRight
	UpLeft
	Right
	Down
	DownRight
	DownLeft
	Left

	velocity int = 5
)

// Move handles direction is pixels the world appears to move. The Player is fixed in the middle of the screen the world
// moves inverse to have the character appear to move in the down direction.
func (d Direction) Move() (x, y int) {
	switch d {
	case Up:
		return 0 * velocity, 1 * velocity
	case UpRight:
		return -1 * velocity, 1 * velocity
	case Right:
		return -1 * velocity, 0 * velocity
	case DownRight:
		return -1 * velocity, -1 * velocity
	case Down:
		return 0 * velocity, -1 * velocity
	case DownLeft:
		return 1 * velocity, -1 * velocity
	case Left:
		return 1 * velocity, 0 * velocity
	case UpLeft:
		return 1 * velocity, 1 * velocity
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

type Input struct {
	mousePosX int
	mousePosY int
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

func (i *Input) MouseState() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (i *Input) MouseLocation() (x, y int) {
	x = i.mousePosX
	y = i.mousePosY

	return
}

func (i *Input) Direction() (Direction, bool) {
	if ebiten.IsKeyPressed(ebiten.KeyW) && ebiten.IsKeyPressed(ebiten.KeyD) {
		return UpRight, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && ebiten.IsKeyPressed(ebiten.KeyS) {
		return DownRight, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && ebiten.IsKeyPressed(ebiten.KeyA) {
		return DownLeft, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) && ebiten.IsKeyPressed(ebiten.KeyW) {
		return UpLeft, true
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
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		return Up, true
	}

	return 0, false
}

package world

import "github.com/hajimehoshi/ebiten"

// Viewport is responsible for drawing everything on the screen.
type Viewport struct {
	screenX int
	screenY int
}

func NewViewport(screenX int, screenY int) *Viewport {
	return &Viewport{
		screenX: screenX,
		screenY: screenY,
	}
}

func Draw(screen *ebiten.Image, offsetX int, offsetY int) {

}

package world

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"reborn/input"
)

type World struct {
	chunks Chunks
	input  *input.Input
	screen *image.Image

	offsetX int
	offsetY int
}

func NewWorld(input *input.Input, screenX int, screenY int) *World {
	w := &World{
		chunks:  make(Chunks),
		input:   input,
		offsetX: -screenX / 2,
		offsetY: -screenY / 2,
	}

	return w
}

func (w *World) OffsetY() int {
	return w.offsetY
}

func (w *World) OffsetX() int {
	return w.offsetX
}

func (w *World) Offset() (x, y int) {
	return w.offsetX, w.offsetY
}

func (w *World) SetOffset(x, y int) {
	w.offsetX = x
	w.offsetY = y
}

func (w *World) MoveOffset(x, y int) {
	w.offsetX += x
	w.offsetY += y
}

func (w *World) MoveWorld(direction input.Direction) {
	w.MoveOffset(direction.Move())
}

// TODO: Need to add to update to add chunks if needed. Ideally padding some space to draw chunks ahead of time. We need to
//		inspect the current offset and based on screen dimensions determine if any chunks need to be populated. For debugging
//		purposes it would be nice to draw chunks and tiles but at least for now at the very least display the chunk key name
//		at mouse cursor location. At the start of the game we should probably initialize chunks ahead of time for the start
//		area. Then the update call only needs to check for chunks at the edges of the screen and avoid checking for chunks
//		if already rendered areas besides game creation.
func (w *World) Update() {

}

func (w *World) Draw(screen *ebiten.Image) {

}

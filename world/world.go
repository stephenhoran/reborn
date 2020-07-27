package world

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/stephenhoran/reborn/debug"
	"github.com/stephenhoran/reborn/input"
	"image"
)

type plane int

const (
	X plane = iota
	Y
)

type World struct {
	chunks Chunks
	input  *input.Input
	screen *image.Image

	offsetX      int
	offsetY      int
	activeChunks []*Chunk

	debugger *debug.Debugger
}

func NewWorld(input *input.Input, screenX int, screenY int, debugger *debug.Debugger) *World {
	w := &World{
		chunks:  make(Chunks),
		input:   input,
		offsetX: screenX / 2,
		offsetY: screenY / 2,

		debugger: debugger,
	}

	w.InitWorld(screenX, screenY)

	return w
}

// Init World renders all of the chunks of the map when first loading the map. Afterward new chunks will be detected and
// added in the Update game loop.
//
// We break up rendering into 4 quadrants of the viewport in which the X and Y numbers will always remain positive or
// negative for simplicity. So far example only things in the top right quadrant of the viewport.
func (w *World) InitWorld(screenX, screenY int) {
	chunkPixels := ChunkSize * TileSize

	// top right quad
	for y := chunkPixels; y < (screenY+w.offsetY)+chunkPixels; y += chunkPixels {
		for x := 0; x < (screenX+w.offsetX)+chunkPixels; x += chunkPixels {
			w.chunks.NewChunk(x, y)
		}
	}

	// bottom right quad
	for y := 0; y > -(screenY+w.offsetY)-chunkPixels; y -= chunkPixels {
		for x := 0; x < (screenX+w.offsetX)+chunkPixels; x += chunkPixels {
			w.chunks.NewChunk(x, y)
		}
	}

	// top left quad
	for y := chunkPixels; y < (screenY+w.offsetY)+chunkPixels; y += chunkPixels {
		for x := -chunkPixels; x > -(screenX+w.offsetX)-chunkPixels; x -= chunkPixels {
			w.chunks.NewChunk(x, y)
		}
	}

	// bottom left quad
	for y := 0; y > -(screenY+w.offsetY)-chunkPixels; y -= chunkPixels {
		for x := -chunkPixels; x > -(screenX+w.offsetX)-chunkPixels; x -= chunkPixels {
			w.chunks.NewChunk(x, y)
		}
	}
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

func (w *World) WorldPositionAtMouse() (int, int) {
	x, y := w.input.MouseLocation()
	return x - w.OffsetX(), w.OffsetY() - y
}

func (w *World) WorldXPositionAtMouse() int {
	return w.input.MouseX() - w.OffsetX()
}

func (w *World) WorldYPositionAtMouse() int {
	return w.OffsetY() - w.input.MouseY()
}

func (w *World) ChunkAtMouse() *Chunk {
	chunk := w.chunks.findChunkAtPosition(w.WorldPositionAtMouse())
	if chunk != nil {
		w.debugger.AddMessage(fmt.Sprintf("Chunk: Chunk_%d_%d", chunk.x, chunk.y))
	}

	return chunk
}

func (w *World) Update() {

}

func (w *World) Draw(screen *ebiten.Image) {
	w.chunks.Draw(screen, w.OffsetX(), w.OffsetY())

	c := w.ChunkAtMouse()
	if c != nil {
		c.DrawChunkTiles(screen, w.offsetX, w.offsetY)
	}

	w.debugger.AddMessage(fmt.Sprintf("X: %d Y: %d", w.input.MouseX(), w.input.MouseY()))
	w.debugger.AddMessage(fmt.Sprintf("World X: %d Y: %d", w.WorldXPositionAtMouse(), w.WorldYPositionAtMouse()))
	w.debugger.AddMessage(fmt.Sprintf("World Offset: X: %d Y: %d", w.OffsetX(), w.OffsetY()))
}

package world

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/stephenhoran/reborn/debug"
	"github.com/stephenhoran/reborn/input"
	"github.com/stephenhoran/reborn/utilities"
)

type Unit int

func (u Unit) Int() int {
	return int(u)
}

const (
	ChunkSize  Unit = 32
	TileSize   Unit = 16
	ChunkPixel      = ChunkSize * TileSize
)

type World struct {
	chunks   Chunks
	input    *input.Input
	viewport *Viewport

	offsetX      int
	offsetY      int
	activeChunks []*Chunk

	debugger *debug.Debugger
}

func NewWorld(input *input.Input, screenX int, screenY int, debugger *debug.Debugger) *World {
	w := &World{
		chunks:   make(Chunks),
		input:    input,
		offsetX:  screenX / 2,
		offsetY:  screenY / 2,
		viewport: NewViewport(screenX, screenY, debugger),

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
	chunkPixels := ChunkPixel.Int()

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
	chunk := w.chunks.FindChunkAtPosition(w.WorldPositionAtMouse())
	if chunk != nil {
		w.debugger.AddMessage(fmt.Sprintf("Chunk: Chunk_%d_%d", chunk.x, chunk.y))
	}

	return chunk
}

func (w *World) TileAtMouse() *Tile {
	tile := w.chunks.FindTileAtPosition(w.WorldPositionAtMouse())
	if tile != nil {
		w.debugger.AddMessage(fmt.Sprintf("Tile X: %d - Y: %d", tile.X(), tile.Y()))
	}

	return tile
}

func (w *World) Update() {
	w.viewport.Update(w.Offset())
}

func (w *World) Draw(screen *ebiten.Image) {
	w.viewport.Draw(screen, w)

	tile := w.TileAtMouse()
	if tile != nil {
		tile.Draw(screen, w.OffsetX(), w.OffsetY())
	}

	w.debugger.AddMessage(fmt.Sprintf("X: %d Y: %d", w.input.MouseX(), w.input.MouseY()))
	w.debugger.AddMessage(fmt.Sprintf("World X: %d Y: %d", w.WorldXPositionAtMouse(), w.WorldYPositionAtMouse()))
	w.debugger.AddMessage(fmt.Sprintf("World Offset: X: %d Y: %d", w.OffsetX(), w.OffsetY()))
}

func findUnit(x, y int, unit Unit) (int, int) {
	unitInt := unit.Int()
	var unitX, unitY int

	if utilities.IsNegativeInt(x) {
		unitX = utilities.Abs(x%unitInt) + x - unitInt
	} else {
		unitX = -(x % unitInt) + x
	}

	if utilities.IsNegativeInt(y) {
		unitY = utilities.Abs(y%unitInt) + y
	} else {
		unitY = -(y % unitInt) + y + unitInt
	}

	return unitX, unitY
}

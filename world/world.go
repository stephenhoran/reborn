package world

import (
	"reborn/input"
	"reborn/utilities"
)

const (
	chunkWidth  int = 100
	chunkHeight int = 100
)

type World struct {
	tiles [][]*Tile
	input *input.Input

	offsetX int
	offsetY int
}

func NewWorld(input *input.Input) *World {
	w := &World{
		input: input,
	}
	w.NewChunk()
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

// NewChunk currently only creates the initial chunk. Until work is done on the camera to allow other chunks.
func (w *World) NewChunk() {
	var x, y int
	tiles := make([][]*Tile, chunkHeight)
	for i := range tiles {
		x = 0
		tw := make([]*Tile, chunkWidth)
		for t := range tw {
			tile := NewTile()
			tile.SetX(x)
			tile.SetY(y)
			tw[t] = tile
			x += TileWidth
		}
		tiles[i] = tw
		y += TileHeight
	}

	w.tiles = tiles
}

// CurrentTile returns the tile at the mouse cursors location.
// The world tiles are kept in relation to the viewport via a world offset.
// TODO: Currently we throw away negative numbers because we have no camera right now.
//		This will need to change as chunks so be able to be rendered in any direction.
func (w *World) CurrentTile() *Tile {
	x, y := w.input.MouseLocation()
	TileX := (x - w.offsetX) / TileWidth
	TileY := (y - w.offsetY) / TileHeight

	if utilities.Abs(TileX) > chunkWidth || utilities.Abs(TileY) > chunkHeight || TileY < 0 || TileX < 0 {
		return NewTile()
	}

	return w.tiles[TileY][TileX]
}

func (w *World) MoveWorld(direction input.Direction) {
	w.MoveOffset(direction.Move())

	for _, tileSlice := range w.tiles {
		for _, tile := range tileSlice {
			tile.Move(direction.Move())
		}
	}
}

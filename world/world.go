package world

import (
	"github.com/hajimehoshi/ebiten"
	"reborn/utilities"
)

const (
	chunkWidth  int = 100
	chunkHeight int = 100
)

type World struct {
	tiles [][]*Tile
}

func NewWorld() *World {
	w := &World{}
	w.NewChunk()
	return w
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
// TODO: Currently we throw away negative numbers because we have no camera right now.
//		This will need to change as chunks so be able to be rendered in any direction.
func (w *World) CurrentTile() *Tile {
	x, y := ebiten.CursorPosition()
	TileX := x / TileWidth
	TileY := y / TileWidth

	if utilities.Abs(TileX) > chunkWidth || utilities.Abs(TileY) > chunkHeight || TileY < 0 || TileX < 0 {
		return NewTile()
	}

	return w.tiles[TileY][TileX]
}

func (w *World) MoveWorld(direction int) {
	switch direction {
	case 0:

	}
}

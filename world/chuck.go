package world

const (
	ChunkSize = 64
)

type Chunks map[string]Chunk

type Chunk struct {
	tiles [][]*Tile

	x int
	y int
}

func (c Chunks) NewChunk(x, y int) {
	chunk := Chunk{
		tiles: make([][]*Tile, ChunkSize),
		x:     x,
		y:     y,
	}

	for i := range chunk.tiles {
		tx := x
		tw := make([]*Tile, ChunkSize)
		for t := range tw {
			tile := NewTile()
			tile.SetX(tx)
			tile.SetY(y)
			tw[t] = tile
			x += TileWidth
		}
		chunk.tiles[i] = tw
		y += TileHeight
	}

	c[string(x)+string(y)] = chunk
}

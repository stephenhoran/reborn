package world

const (
	TileWidth  int = 32
	TileHeight int = 32
)

type Tile struct {
	x int
	y int
}

func NewTile() *Tile {
	return &Tile{}
}

func (t *Tile) X() int {
	return t.x
}

func (t *Tile) SetX(x int) {
	t.x = x
}

func (t *Tile) Y() int {
	return t.y
}

func (t *Tile) SetY(y int) {
	t.y = y
}

func (t *Tile) Move(x, y int) {
	t.x += x
	t.y += y
}

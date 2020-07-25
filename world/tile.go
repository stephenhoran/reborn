package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

const (
	TileWidth  int = 32
	TileHeight int = 32
)

type Tile struct {
	x int
	y int

	draw bool
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

func (t *Tile) DrawStatus() bool {
	return t.draw
}

func (t *Tile) SetDrawStatus(b bool) {
	t.draw = b
}

func (t *Tile) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, float64(t.X()), float64(t.Y()), float64(TileWidth), float64(TileHeight), color.RGBA{R: 114, G: 127, B: 140, A: 255})
}

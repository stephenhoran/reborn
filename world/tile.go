package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

const (
	TileSize int = 16
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

func (t *Tile) Draw(screen *ebiten.Image, offsetX int, offsetY int) {
	x := t.X()
	y := t.Y()

	tileColor := color.RGBA{R: 81, G: 188, B: 255, A: 255}

	// Top Line - Left Line - Bottom Line - Right Line

	ebitenutil.DrawLine(screen, float64(x+offsetX), float64(offsetY-y), float64(x+offsetX+TileSize), float64(offsetY-y), tileColor)
	ebitenutil.DrawLine(screen, float64(x+offsetX), float64(offsetY-y), float64(x+offsetX), float64(offsetY-y-TileSize), tileColor)
	ebitenutil.DrawLine(screen, float64(x+offsetX), float64(offsetY-y-TileSize), float64(x+offsetX+TileSize), float64(offsetY-y-TileSize), tileColor)
	ebitenutil.DrawLine(screen, float64(x+offsetX+TileSize), float64(offsetY-y), float64(x+offsetX+TileSize), float64(offsetY-y-TileSize), tileColor)

}

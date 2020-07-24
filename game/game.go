package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"reborn/assets"
	"reborn/world"
)

const (
	screenWidth  int = 1280
	screenHeight int = 720
)

type Game struct {
	input *Input
	world *world.World
}

var asset assets.Assets

func init() {
	ebiten.SetWindowTitle("Reborn")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	asset = assets.LoadAssets()
}

func NewGame() *Game {
	return &Game{
		input: NewInput(),
		world: world.NewWorld(),
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.input.Update()
	dir, ok := g.input.Direction()
	if ok {
		log.Println(dir.String())
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 135, G: 211, B: 124, A: 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(50, 50)
	op.GeoM.Scale(.5, .5)
	screen.DrawImage(asset["player"].SubImage(image.Rect(0, 0, 80, 85)).(*ebiten.Image), op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %d Y: %d", g.input.mousePosX, g.input.mousePosY))
	tile := g.world.CurrentTile()
	ebitenutil.DrawRect(screen, float64(tile.X()), float64(tile.Y()), float64(world.TileWidth), float64(world.TileHeight), color.RGBA{R: 114, G: 127, B: 140, A: 255})
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Current Tile: X: %d Y: %d", tile.X(), tile.Y()), 0, 15)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

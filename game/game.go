package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"image/color"
	_ "image/png"
	"reborn/assets"
	"reborn/input"
	"reborn/world"
)

const (
	screenWidth  int = 1280
	screenHeight int = 720
)

type Game struct {
	input *input.Input
	world *world.World
}

var asset assets.Assets

func init() {
	ebiten.SetWindowTitle("Reborn")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	asset = assets.LoadAssets()
}

func NewGame() *Game {
	i := input.NewInput()

	return &Game{
		input: i,
		world: world.NewWorld(i, screenWidth, screenHeight),
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.input.Update()
	dir, ok := g.input.Direction()
	if ok {
		g.world.MoveWorld(dir)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 135, G: 211, B: 124, A: 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(screenWidth), float64(screenHeight))
	op.GeoM.Scale(.5, .5)
	screen.DrawImage(asset["player"].SubImage(image.Rect(0, 0, 80, 80)).(*ebiten.Image), op)

	g.world.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %d Y: %d", g.input.MouseX(), g.input.MouseY()))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("World Offset: X: %d Y: %d", g.world.OffsetX(), g.world.OffsetY()), 0, 30)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

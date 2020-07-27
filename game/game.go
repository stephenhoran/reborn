package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/stephenhoran/reborn/assets"
	"github.com/stephenhoran/reborn/debug"
	"github.com/stephenhoran/reborn/input"
	"github.com/stephenhoran/reborn/world"
	"image"
	"image/color"
	_ "image/png"
)

const (
	screenWidth  int = 1280
	screenHeight int = 720
)

type Game struct {
	input *input.Input
	world *world.World

	debugger *debug.Debugger
}

var asset assets.Assets

func init() {
	ebiten.SetWindowTitle("Reborn")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	asset = assets.LoadAssets()
}

func NewGame() *Game {
	i := input.NewInput()

	debugger := debug.NewDebugger()

	return &Game{
		input:    i,
		world:    world.NewWorld(i, screenWidth, screenHeight, debugger),
		debugger: debugger,
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.input.Update()
	g.world.Update()
	dir, ok := g.input.Direction()
	if ok {
		g.world.MoveWorld(dir)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 135, G: 211, B: 124, A: 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(screenWidth-40), float64(screenHeight-40))
	op.GeoM.Scale(.5, .5)
	screen.DrawImage(asset["player"].SubImage(image.Rect(0, 0, 80, 80)).(*ebiten.Image), op)

	g.world.Draw(screen)

	g.debugger.AddMessage(fmt.Sprintf("Current FPS: %f", ebiten.CurrentFPS()))
	g.debugger.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

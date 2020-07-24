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
)

const (
	screenWidth  int = 800
	screenHeight int = 600
)

type Game struct {
	input *Input
}

var asset assets.Assets

func init() {
	ebiten.SetWindowTitle("Reborn")
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	asset = assets.LoadAssets()
}

func NewGame() *Game {
	return &Game{
		input: NewInput(),
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
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

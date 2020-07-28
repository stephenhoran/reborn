package world

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/stephenhoran/reborn/debug"
)

type Coordinates struct {
	X int
	Y int
}

// Viewport is responsible for drawing everything on the screen.
type Viewport struct {
	screenX int
	screenY int

	offsetX int
	offsetY int

	worldPositionX int
	worldPositionY int

	topLeft     Coordinates
	topRight    Coordinates
	bottomLeft  Coordinates
	bottomRight Coordinates

	debugger *debug.Debugger
}

func (v *Viewport) OffsetY() int {
	return v.offsetY
}

func (v *Viewport) SetOffsetY(offsetY int) {
	v.offsetY = offsetY
}

func (v *Viewport) OffsetX() int {
	return v.offsetX
}

func (v *Viewport) SetOffsetX(offsetX int) {
	v.offsetX = offsetX
}

func (v *Viewport) SetWorldPosition(offsetX, offsetY int) {
	v.SetWorldPositionX(v.ScreenX()/2 - offsetX)
	v.SetWorldPositionY(offsetY - v.ScreenY()/2)

	v.topLeft.X = v.WorldPositionX() - v.ScreenX()/2
	v.topLeft.Y = v.WorldPositionY() + v.ScreenY()/2

	v.topRight.X = v.WorldPositionX() + v.ScreenX()/2
	v.topRight.Y = v.WorldPositionY() + v.ScreenY()/2

	v.bottomLeft.X = v.WorldPositionX() - v.ScreenX()/2
	v.bottomLeft.Y = v.WorldPositionY() - v.ScreenY()/2

	v.bottomRight.X = v.WorldPositionX() + v.ScreenX()/2
	v.bottomRight.Y = v.WorldPositionY() - v.ScreenY()/2
}

func (v *Viewport) WorldPositionY() int {
	return v.worldPositionY
}

func (v *Viewport) SetWorldPositionY(worldPositionY int) {
	v.worldPositionY = worldPositionY
}

func (v *Viewport) WorldPositionX() int {
	return v.worldPositionX
}

func (v *Viewport) SetWorldPositionX(worldPositionX int) {
	v.worldPositionX = worldPositionX
}

func (v *Viewport) ScreenY() int {
	return v.screenY
}

func (v *Viewport) SetScreenY(screenY int) {
	v.screenY = screenY
}

func (v *Viewport) ScreenX() int {
	return v.screenX
}

func (v *Viewport) SetScreenX(screenX int) {
	v.screenX = screenX
}

func NewViewport(screenX int, screenY int, debugger *debug.Debugger) *Viewport {
	return &Viewport{
		screenX:  screenX,
		screenY:  screenY,
		debugger: debugger,

		topLeft:     Coordinates{},
		topRight:    Coordinates{},
		bottomLeft:  Coordinates{},
		bottomRight: Coordinates{},
	}
}

func (v *Viewport) Update(offsetX, offsetY int) {
	v.SetWorldPosition(offsetX, offsetY)
}

func (v *Viewport) Draw(screen *ebiten.Image, world *World) {
	v.debugger.AddMessage(fmt.Sprintf("Viewport X: %d - Y: %d", v.WorldPositionX(), v.WorldPositionY()))
	v.debugger.AddMessage(fmt.Sprintf("Top Left X: %d - Y: %d", v.topLeft.X, v.topLeft.Y))
	v.debugger.AddMessage(fmt.Sprintf("Top Right X: %d - Y: %d", v.topRight.X, v.topRight.Y))
	v.debugger.AddMessage(fmt.Sprintf("Bottom Left X: %d - Y: %d", v.bottomLeft.X, v.bottomLeft.Y))
	v.debugger.AddMessage(fmt.Sprintf("Bottom Right X: %d - Y: %d", v.bottomRight.X, v.bottomRight.Y))

	v.DrawScreenTiles(screen, world)
}

func (v *Viewport) DrawScreenTiles(screen *ebiten.Image, world *World) {
	for y := v.topLeft.Y; y > v.bottomLeft.Y-TileSize.Int(); y -= TileSize.Int() {
		for x := v.topLeft.X; x < v.topRight.X+TileSize.Int(); x += TileSize.Int() {
			tile := world.chunks.FindTileAtPosition(x, y)
			tile.Draw(screen, world.OffsetX(), world.OffsetY())
		}
	}
}

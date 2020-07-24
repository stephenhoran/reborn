package main

import (
	"github.com/hajimehoshi/ebiten"
	"reborn/game"
)

func main() {
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

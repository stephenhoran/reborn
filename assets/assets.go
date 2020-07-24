package assets

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Assets map[string]*ebiten.Image

type Location string

func LoadAssets() Assets {
	img, _, err := ebitenutil.NewImageFromFile("./assets/character/player.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}

	return Assets{
		"player": img,
	}
}

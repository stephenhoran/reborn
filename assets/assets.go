package assets

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/png"
	"reborn/assets/character"
)

type Assets map[string]*ebiten.Image

type Location string

func LoadAssets() Assets {
	img, _, _ := image.Decode(bytes.NewBuffer(character.Player_png))

	player, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}

	return Assets{
		"player": player,
	}
}

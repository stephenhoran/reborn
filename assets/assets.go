package assets

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"github.com/stephenhoran/reborn/assets/character"
	"image"
	_ "image/png"
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

package wrapper

import (
	"errors"

	"github.com/telroshan/go-sfml/v2/graphics"
)

const cTrue = 1
const cFalse = 0

type Resources struct {
	textures []graphics.Struct_SS_sfTexture
	sprites  []Sprite
}

func (res *Resources) AddTexture(item graphics.Struct_SS_sfTexture) {
	res.textures = append(res.textures, item)
}

func (res *Resources) AddSprite(item graphics.Struct_SS_sfSprite) {
	res.sprites = append(res.sprites, Sprite{item})
}

func (res *Resources) Clear() {
	for _, item := range res.textures {
		graphics.SfTexture_destroy(item)
	}
	res.textures = nil
	for _, item := range res.sprites {
		graphics.SfSprite_destroy(item.sprite)
	}
	res.sprites = nil
}

func FileToSprite(filename string, res *Resources) (*Sprite, error) {
	t := graphics.SfTexture_createFromFile(filename, getNullIntRect())
	if t == nil || t.Swigcptr() == cFalse {
		return nil, errors.New("Couldn't load png")
	}
	res.AddTexture(t)
	s := graphics.SfSprite_create()
	graphics.SfSprite_setTexture(s, t, cTrue)
	res.AddSprite(s)
	return &res.sprites[len(res.sprites)-1], nil
}

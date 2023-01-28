package wrapper

import "github.com/telroshan/go-sfml/v2/graphics"

type Sprite struct {
	sprite graphics.Struct_SS_sfSprite
}

func (sprite *Sprite) Move(x float32, y float32) {
	graphics.SfShape_move(sprite.sprite, makeVector2(x, y))
}

func (sprite *Sprite) SetTextureRect(left int, top int, width int, height int) {
	rect := graphics.NewSfIntRect()
	rect.SetLeft(left)
	rect.SetTop(top)
	rect.SetWidth(width)
	rect.SetHeight(height)
	graphics.SfSprite_setTextureRect(sprite.sprite, rect)
}

func (sprite *Sprite) SetPosition(x float32, y float32) {
	graphics.SfRectangleShape_setPosition(sprite.sprite, makeVector2(x, y))
}

func (sprite *Sprite) Draw(w graphics.Struct_SS_sfRenderWindow) {
	graphics.SfRenderWindow_drawSprite(w, sprite.sprite, getNullRenderState())
}

func makeVector2(x float32, y float32) graphics.SfVector2f {
	v := graphics.NewSfVector2f()
	v.SetX(x)
	v.SetY(y)
	return v
}

func getNullIntRect() graphics.SfIntRect {
	return (graphics.SfIntRect)(graphics.SwigcptrSfIntRect(0))
}

func getNullRenderState() graphics.SfRenderStates {
	return (graphics.SfRenderStates)(graphics.SwigcptrSfRenderStates(0))
}

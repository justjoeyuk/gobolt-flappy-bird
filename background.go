package main

import "github.com/justjoeyuk/gobolt/display"

// Background -
type Background struct {
	*display.Sprite
}

// NewBackground -
func NewBackground() *Background {
	var err error
	bg := &Background{Sprite: display.NewSprite()}

	bg.Width = 1920
	bg.Height = 1080
	bg.Y = -420
	bg.Texture, err = display.LoadTextureFromFile("res/images/background.png")

	if err != nil {
		panic("Could not create BG")
	}

	return bg
}

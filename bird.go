package main

import "github.com/justjoeyuk/gobolt/display"

// Bird -
type Bird struct {
	*display.Sprite

	FallSpeed float32
}

// NewBird -
func NewBird() *Bird {
	var err error
	bird := &Bird{Sprite: display.NewSprite()}

	bird.FallSpeed = 0
	bird.Width = 50
	bird.Height = 43
	bird.X = 10
	bird.Y = 300 - (bird.Height / 2)
	bird.Texture, err = display.LoadTextureFromFile("res/images/bird.png")

	if err != nil {
		panic("Could not create Bird")
	}

	return bird
}

// Update -
func (b *Bird) Update() {
	b.Y += int32(b.FallSpeed)
	b.FallSpeed += 0.4

	if b.Y > 600-b.Height {
		b.FallSpeed = -10
	}
}

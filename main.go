package main

import (
	"os"

	"github.com/justjoeyuk/gobolt/core"
	"github.com/veandco/go-sdl2/sdl"
)

// Engine - GOBOLT
var Engine *core.Engine

func main() {
	sdl.Main(func() {
		Engine := &core.Engine{
			ResX:        600,
			ResY:        600,
			FPS:         60,
			WindowTitle: "My Game",
		}

		Engine.Initialize()

		if err := Engine.Run(NewWorld()); err != nil {
			panic(err)
		}

		os.Exit(0)
	})
}

package main

import (
	"github.com/justjoeyuk/gobolt/core"
	"github.com/veandco/go-sdl2/sdl"
)

// World - The Game Object
type World struct {
	*core.Game

	bg   *Background
	bird *Bird
}

// NewWorld -
func NewWorld() *World {
	return &World{Game: core.NewGame()}
}

// Start - Begin the Game
func (w *World) Start() {
	w.bird = NewBird()
	w.bg = NewBackground()

	w.AddChild(w.bg.Sprite)
	w.AddChild(w.bird.Sprite)
}

// Update - Update the Game
func (w *World) Update() {
	w.bird.Update()
}

// HandleEvent - Handle any events
func (w *World) HandleEvent(e sdl.Event) {
	//fmt.Printf("Event: %T", e)
}

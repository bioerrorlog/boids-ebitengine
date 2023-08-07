package main

import (
	"log"

	"github.com/bioerrorlog/boids-ebitengine/boids"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := boids.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

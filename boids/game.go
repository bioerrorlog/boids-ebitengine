package boids

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

type Game struct {
	boids []*Boid
}

func NewGame() (*Game, error) {
	g := &Game{}
	return g, nil
}

func (g *Game) Update() error {
	for _, b := range g.boids {
		b.Update(g.boids)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Backgroud
	screen.Fill(color.RGBA{15, 2, 22, 0xff})

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

package boids

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/bioerrorlog/boids-ebitengine/vector"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
	boidCount    = 100
)

type Game struct {
	boids []*Boid
}

func NewGame() (*Game, error) {
	g := &Game{
		boids: make([]*Boid, boidCount),
	}

	for i := range g.boids {
		g.boids[i] = NewBoid(
			rand.Float64()*ScreenWidth,
			rand.Float64()*ScreenHeight,
			vector.Vec2{X: ScreenWidth / 2, Y: ScreenHeight / 2},
		)
	}
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

	// Boids
	for _, b := range g.boids {
		b.Draw(screen)
	}

	// Debug
	fps := fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, fps)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

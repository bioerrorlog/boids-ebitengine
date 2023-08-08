package boids

import (
	"image/color"
	"math/rand"

	"github.com/bioerrorlog/boids-ebitengine/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

type Boid struct {
	position, velocity vector.Vec2
}

func NewBoid(x, y float64) *Boid {
	return &Boid{
		position: vector.Vec2{X: x, Y: y},
		velocity: vector.Vec2{X: rand.Float64()*2 - 1, Y: rand.Float64()*2 - 1},
	}
}

func (b *Boid) Draw(screen *ebiten.Image) {
	boidImage := ebiten.NewImage(20, 20)
	boidImage.Fill(color.White)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(b.position.X, b.position.Y)

	screen.DrawImage(boidImage, opts)
}

func (b *Boid) Update(boids []*Boid) {
	alignment := b.alignment(boids)
	cohesion := b.cohesion(boids)
	separation := b.separation(boids)

	b.velocity = b.velocity.Add(alignment).Add(cohesion).Add(separation).Limit(2)
	b.position = b.position.Add(b.velocity)
}

func (b *Boid) alignment(boids []*Boid) vector.Vec2 {
	var sum vector.Vec2
	var count int
	for _, other := range boids {
		if b != other {
			sum = sum.Add(other.velocity)
			count++
		}
	}

	if count > 0 {
		avg := sum.Div(float64(count))
		return avg.Sub(b.velocity).Limit(0.05)
	}

	return vector.Vec2{}
}

func (b *Boid) cohesion(boids []*Boid) vector.Vec2 {
	var sum vector.Vec2
	var count int
	for _, other := range boids {
		if b != other {
			sum = sum.Add(other.position)
			count++
		}
	}

	if count > 0 {
		avg := sum.Div(float64(count))
		return avg.Sub(b.position).Limit(0.05)
	}

	return vector.Vec2{}
}

func (b *Boid) separation(boids []*Boid) vector.Vec2 {
	var sum vector.Vec2
	for _, other := range boids {
		if b != other {
			diff := b.position.Sub(other.position)
			if diff.Length() < 30 {
				sum = sum.Add(diff.Normalize().Div(diff.Length()))
			}
		}
	}

	return sum.Limit(0.5)
}

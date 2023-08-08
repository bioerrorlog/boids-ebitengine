package boids

import (
	"math/rand"

	"github.com/bioerrorlog/boids-ebitengine/vector"
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

func (b *Boid) Update(boids []*Boid) {
	alignment := b.alignment(boids)
	cohesion := b.cohesion(boids)
	separation := b.separation(boids)

	b.velocity = b.velocity.Add(alignment).Add(cohesion).Add(separation).Limit(2)
	b.position = b.position.Add(b.velocity)
}

func (b *Boid) alignment(boids []*Boid) vector.Vec2 {
	// TODO
	return vector.Vec2{X: 0, Y: 0}
}

func (b *Boid) cohesion(boids []*Boid) vector.Vec2 {
	// TODO
	return vector.Vec2{X: 0, Y: 0}
}

func (b *Boid) separation(boids []*Boid) vector.Vec2 {
	// TODO
	return vector.Vec2{X: 0, Y: 0}
}

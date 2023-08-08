package boids

import "github.com/bioerrorlog/boids-ebitengine/vector"

type Boid struct {
	position, velocity vector.Vec2
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
	return vector.Vec2{0, 0}
}

func (b *Boid) cohesion(boids []*Boid) vector.Vec2 {
	// TODO
	return vector.Vec2{0, 0}
}

func (b *Boid) separation(boids []*Boid) vector.Vec2 {
	// TODO
	return vector.Vec2{0, 0}
}

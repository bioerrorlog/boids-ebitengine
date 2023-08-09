package boids

import (
	"image/color"
	"math/rand"

	"github.com/bioerrorlog/boids-ebitengine/vector"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	moveSpeed                 = 20
	perceptionRadius          = 100
	steerForce                = 1
	alignmentForce            = 0.1
	cohesionForce             = 0.1
	separationForce           = 0.1
	centralizationForce       = 0.4
	centralizationForceRadius = 200
)

type Boid struct {
	position, velocity, targetCenter vector.Vec2
}

func NewBoid(x, y float64, targetCenter vector.Vec2) *Boid {
	return &Boid{
		position:     vector.Vec2{X: x, Y: y},
		velocity:     vector.Vec2{X: rand.Float64()*2 - 1, Y: rand.Float64()*2 - 1},
		targetCenter: targetCenter,
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
	neighbors := b.getNeighbors(boids)

	alignment := b.alignment(neighbors)
	cohesion := b.cohesion(neighbors)
	separation := b.separation(neighbors)
	centering := b.centralization()

	b.velocity = b.velocity.Add(alignment).Add(cohesion).Add(separation).Add(centering).Limit(moveSpeed)
	b.position = b.position.Add(b.velocity)
}

func (b *Boid) SetTargetCenter(center vector.Vec2) {
	b.targetCenter = center
}

func (b *Boid) alignment(boids []*Boid) vector.Vec2 {
	var sum vector.Vec2
	if len(boids) == 0 {
		return sum
	}
	for _, other := range boids {
		if b != other {
			sum = sum.Add(other.velocity)
		}
	}
	avg := sum.Div(float64(len(boids)))
	return b.steer(avg.Normalize().Mul(moveSpeed)).Mul(alignmentForce)
}

func (b *Boid) cohesion(boids []*Boid) vector.Vec2 {
	var sum vector.Vec2
	if len(boids) == 0 {
		return sum
	}
	for _, other := range boids {
		if b != other {
			sum = sum.Add(other.position)
		}
	}
	avg := sum.Div(float64(len(boids)))
	return b.steer(avg.Sub(b.position).Normalize().Mul(moveSpeed)).Mul(cohesionForce)
}

func (b *Boid) separation(boids []*Boid) vector.Vec2 {
	var sum vector.Vec2
	var closeNeighbors []*Boid
	for _, other := range boids {
		if b != other && b.position.DistanceTo(other.position) < perceptionRadius/2 {
			closeNeighbors = append(closeNeighbors, other)
		}
	}
	if len(closeNeighbors) == 0 {
		return sum
	}
	for _, other := range closeNeighbors {
		diff := b.position.Sub(other.position)
		sum = sum.Add(diff.Normalize().Div(diff.Length()))
	}
	avg := sum.Div(float64(len(closeNeighbors)))
	return b.steer(avg.Normalize().Mul(moveSpeed)).Mul(separationForce)
}

func (b *Boid) centralization() vector.Vec2 {
	if b.position.DistanceTo(b.targetCenter) < centralizationForceRadius {
		return vector.Vec2{}
	}
	desired := b.targetCenter.Sub(b.position).Normalize().Mul(moveSpeed)
	return b.steer(desired).Mul(centralizationForce)
}

func (b *Boid) steer(target vector.Vec2) vector.Vec2 {
	steer := target.Sub(b.velocity)
	return steer.Normalize().Mul(steerForce)
}
func (b *Boid) getNeighbors(boids []*Boid) []*Boid {
	var neighbors []*Boid
	for _, other := range boids {
		if b != other && b.position.DistanceTo(other.position) < perceptionRadius {
			neighbors = append(neighbors, other)
		}
	}
	return neighbors
}

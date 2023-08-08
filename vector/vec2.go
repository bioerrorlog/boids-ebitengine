package vector

import "math"

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

func (v Vec2) Mul(scalar float64) Vec2 {
	return Vec2{v.X * scalar, v.Y * scalar}
}

func (v Vec2) Div(scalar float64) Vec2 {
	return Vec2{v.X / scalar, v.Y / scalar}
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) Normalize() Vec2 {
	length := v.Length()
	if length != 0 {
		return Vec2{v.X / length, v.Y / length}
	}
	return Vec2{0, 0}
}

func (v Vec2) limit(max float64) Vec2 {
	if v.Length() > max {
		return v.Normalize().Mul(max)
	}
	return v
}

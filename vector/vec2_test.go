package vector

import (
	"testing"
)

func TestVec2_Add(t *testing.T) {
	tests := []struct {
		name  string
		v     Vec2
		other Vec2
		want  Vec2
	}{
		{"basic add", Vec2{3, 4}, Vec2{1, 2}, Vec2{4, 6}},
		{"add with zero", Vec2{3, 4}, Vec2{0, 0}, Vec2{3, 4}},
		{"add with negative values", Vec2{3, 4}, Vec2{-1, -2}, Vec2{2, 2}},
		{"add with same values", Vec2{3, 4}, Vec2{3, 4}, Vec2{6, 8}},
		{"add with floating values", Vec2{3.5, 4.5}, Vec2{1.5, 2.5}, Vec2{5, 7}},
		{"add with mix of positive and negative", Vec2{3, -4}, Vec2{-1, 2}, Vec2{2, -2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Add(tt.other); got != tt.want {
				t.Errorf("Vec2.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Sub(t *testing.T) {
	tests := []struct {
		name  string
		v     Vec2
		other Vec2
		want  Vec2
	}{
		{"basic subtraction", Vec2{3, 4}, Vec2{1, 2}, Vec2{2, 2}},
		{"subtract with zero", Vec2{3, 4}, Vec2{0, 0}, Vec2{3, 4}},
		{"subtract negative values", Vec2{3, 4}, Vec2{-1, -2}, Vec2{4, 6}},
		{"subtract same values", Vec2{3, 4}, Vec2{3, 4}, Vec2{0, 0}},
		{"subtract floating values", Vec2{3.5, 4.5}, Vec2{1.5, 2.5}, Vec2{2, 2}},
		{"subtract mix of positive and negative", Vec2{3, -4}, Vec2{-1, 2}, Vec2{4, -6}},
		{"subtract resulting in negative", Vec2{3, 4}, Vec2{5, 6}, Vec2{-2, -2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Sub(tt.other); got != tt.want {
				t.Errorf("Vec2.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Mul(t *testing.T) {
	tests := []struct {
		name   string
		v      Vec2
		scalar float64
		want   Vec2
	}{
		{"basic multiplication", Vec2{3, 4}, 2, Vec2{6, 8}},
		{"multiply by zero", Vec2{3, 4}, 0, Vec2{0, 0}},
		{"multiply by one", Vec2{3, 4}, 1, Vec2{3, 4}},
		{"multiply by negative", Vec2{3, 4}, -2, Vec2{-6, -8}},
		{"multiply by fractional", Vec2{3, 4}, 0.5, Vec2{1.5, 2}},
		{"multiply with negative and fractional", Vec2{-3, 4}, -0.5, Vec2{1.5, -2}},
		{"multiply with floating values", Vec2{3.5, 4.5}, 1.5, Vec2{5.25, 6.75}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Mul(tt.scalar); got != tt.want {
				t.Errorf("Vec2.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Div(t *testing.T) {
	tests := []struct {
		name   string
		v      Vec2
		scalar float64
		want   Vec2
	}{
		{"basic division", Vec2{6, 8}, 2, Vec2{3, 4}},
		{"divide by one", Vec2{3, 4}, 1, Vec2{3, 4}},
		{"divide by negative", Vec2{-6, 8}, -2, Vec2{3, -4}},
		{"divide by fractional", Vec2{3, 4}, 0.5, Vec2{6, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Div(tt.scalar); got != tt.want {
				t.Errorf("Vec2.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Length(t *testing.T) {
	tests := []struct {
		name string
		v    Vec2
		want float64
	}{
		{"basic length", Vec2{3, 4}, 5},
		{"zero length", Vec2{0, 0}, 0},
		{"negative values", Vec2{-3, 4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Length(); got != tt.want {
				t.Errorf("Vec2.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Normalize(t *testing.T) {
	tests := []struct {
		name string
		v    Vec2
		want Vec2
	}{
		{"basic normalize", Vec2{3, 4}, Vec2{3 / 5.0, 4 / 5.0}},
		{"normalize zero vector", Vec2{0, 0}, Vec2{0, 0}},
		{"normalize negative values", Vec2{-3, 4}, Vec2{-3 / 5.0, 4 / 5.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Normalize(); got != tt.want {
				t.Errorf("Vec2.Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Limit(t *testing.T) {
	tests := []struct {
		name string
		v    Vec2
		max  float64
		want Vec2
	}{
		{"limit within length", Vec2{3, 4}, 5, Vec2{3, 4}},
		{"limit to length", Vec2{3, 4}, 3, Vec2{3 / 5.0 * 3, 4 / 5.0 * 3}},
		{"limit with zero vector", Vec2{0, 0}, 5, Vec2{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Limit(tt.max); got != tt.want {
				t.Errorf("Vec2.Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_DistanceTo(t *testing.T) {
	tests := []struct {
		name  string
		v     Vec2
		other Vec2
		want  float64
	}{
		{"basic distance", Vec2{3, 4}, Vec2{0, 0}, 5},
		{"distance to self", Vec2{3, 4}, Vec2{3, 4}, 0},
		{"distance with negative values", Vec2{3, 4}, Vec2{-3, 4}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.DistanceTo(tt.other); got != tt.want {
				t.Errorf("Vec2.DistanceTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

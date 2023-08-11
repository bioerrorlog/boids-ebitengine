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

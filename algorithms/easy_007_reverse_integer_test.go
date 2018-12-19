package algorithms

import "testing"

func TestReverseInteger(t *testing.T) {
	cases := []struct {
		x    int
		want int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1, 1},
		{100, 1},
		{1534236469, 0},
	}
	for _, c := range cases {
		got := reverse(c.x)
		if c.want != got {
			t.Errorf("reverse(%v) == %v, want %v", c.x, got, c.want)
		}
	}
}

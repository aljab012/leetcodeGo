package main

import (
	"strconv"
	"testing"
)

func TestTribonacci(t *testing.T) {
	testCases := []struct {
		n, want int
	}{
		{4, 4},
		{25, 1389537},
	}

	funcs := []struct {
		name string
		fn   func(int) int
	}{
		{"tribonacci1", tribonacci1},
		{"tribonacci2", tribonacci2},
		{"tribonacci3", tribonacci3},
		{"tribonacci4", tribonacci4},
	}

	for _, f := range funcs {
		f := f
		t.Run(f.name, func(t *testing.T) {
			for _, tc := range testCases {
				tc := tc
				t.Run("n="+strconv.Itoa(tc.n), func(t *testing.T) {
					if got := f.fn(tc.n); got != tc.want {
						t.Errorf("for n=%d: got %d, want %d", tc.n, got, tc.want)
					}
				})
			}
		})
	}
}

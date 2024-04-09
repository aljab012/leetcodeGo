package main

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"Example 2", args{height: []int{4, 2, 0, 3, 2, 5}}, 9},
		{"Example 3", args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"Example 4", args{height: []int{4, 2, 0, 3, 2, 5}}, 9},
		{"Example 5", args{height: []int{4, 2, 3}}, 1},
		{"Example 6", args{height: []int{4, 2, 3, 1}}, 1},
		{"Example 7", args{height: []int{4, 2, 3, 1, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{x: 1, y: 2}, 1},
		{"Example 2", args{x: 2, y: 1}, 1},
		{"Example 3", args{x: 1, y: 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{x: 1, y: 2}, 2},
		{"Example 2", args{x: 2, y: 1}, 2},
		{"Example 3", args{x: 1, y: 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

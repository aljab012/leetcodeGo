package main

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"Example 2", args{height: []int{1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
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
		{"Example 4", args{x: 2, y: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
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
		{"Example 4", args{x: 2, y: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

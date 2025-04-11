package main

import (
	"testing"
)

func Test_minPathSum1(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}}, 7},
		{"Test 2", args{[][]int{{1, 2, 3}, {4, 5, 6}}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum1(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum1() = %v, want %v", got, tt.want)
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
		{"Test 1", args{1, 2}, 1},
		{"Test 2", args{2, 1}, 1},
		{"Test 3", args{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPathSum2(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}}, 7},
		{"Test 2", args{[][]int{{1, 2, 3}, {4, 5, 6}}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum2(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

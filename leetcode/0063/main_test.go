package main

import "testing"

func Test_uniquePathsWithObstacles1(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, 2},
		{"Test 2", args{[][]int{{0, 1}, {0, 0}}}, 1},
		{"Test 3", args{[][]int{{0, 0}, {1, 1}, {0, 0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles1(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePathsWithObstacles2(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, 2},
		{"Test 2", args{[][]int{{0, 1}, {0, 0}}}, 1},
		{"Test 3", args{[][]int{{0, 0}, {1, 1}, {0, 0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles2(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles2() = %v, want %v", got, tt.want)
			}
		})
	}
}

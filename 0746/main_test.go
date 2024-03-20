package main

import "testing"

func Test_minCostClimbingStairs1(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]int{10, 15, 20}}, 15},
		{"Test 2", args{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, 6},
		{"Test 3", args{[]int{0, 0, 0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs1(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCostClimbingStairs2(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]int{10, 15, 20}}, 15},
		{"Test 2", args{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, 6},
		{"Test 3", args{[]int{0, 0, 0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs2(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCostClimbingStairs3(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]int{10, 15, 20}}, 15},
		{"Test 2", args{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, 6},
		{"Test 3", args{[]int{0, 0, 0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs3(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCostClimbingStairs4(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]int{10, 15, 20}}, 15},
		{"Test 2", args{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, 6},
		{"Test 3", args{[]int{0, 0, 0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs4(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs4() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

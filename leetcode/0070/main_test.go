package main

import "testing"

func Test_climbStairs1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{2}, 2},
		{"Test 2", args{3}, 3},
		{"Test 3", args{4}, 5},
		{"Test 4", args{5}, 8},
		{"Test 5", args{6}, 13},
		{"Test 6", args{7}, 21},
		{"Test 7", args{8}, 34},
		{"Test 8", args{9}, 55},
		{"Test 9", args{10}, 89},
		{"Test 10", args{11}, 144},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs1(tt.args.n); got != tt.want {
				t.Errorf("climbStairs1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{2}, 2},
		{"Test 2", args{3}, 3},
		{"Test 3", args{4}, 5},
		{"Test 4", args{5}, 8},
		{"Test 5", args{6}, 13},
		{"Test 6", args{7}, 21},
		{"Test 7", args{8}, 34},
		{"Test 8", args{9}, 55},
		{"Test 9", args{10}, 89},
		{"Test 10", args{11}, 144},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs2(tt.args.n); got != tt.want {
				t.Errorf("climbStairs2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{2}, 2},
		{"Test 2", args{3}, 3},
		{"Test 3", args{4}, 5},
		{"Test 4", args{5}, 8},
		{"Test 5", args{6}, 13},
		{"Test 6", args{7}, 21},
		{"Test 7", args{8}, 34},
		{"Test 8", args{9}, 55},
		{"Test 9", args{10}, 89},
		{"Test 10", args{11}, 144},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs3(tt.args.n); got != tt.want {
				t.Errorf("climbStairs3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs4(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{2}, 2},
		{"Test 2", args{3}, 3},
		{"Test 3", args{4}, 5},
		{"Test 4", args{5}, 8},
		{"Test 5", args{6}, 13},
		{"Test 6", args{7}, 21},
		{"Test 7", args{8}, 34},
		{"Test 8", args{9}, 55},
		{"Test 9", args{10}, 89},
		{"Test 10", args{11}, 144},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs4(tt.args.n); got != tt.want {
				t.Errorf("climbStairs4() = %v, want %v", got, tt.want)
			}
		})
	}
}

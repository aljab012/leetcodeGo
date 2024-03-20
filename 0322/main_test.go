package main

import "testing"

func Test_coinChange1(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]int{1, 2, 5}, 11}, 3},
		{"Test 2", args{[]int{2}, 3}, -1},
		{"Test 3", args{[]int{1}, 0}, 0},
		{"Test 4", args{[]int{1}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange1(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChange2(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]int{1, 2, 5}, 11}, 3},
		{"Test 2", args{[]int{2}, 3}, -1},
		{"Test 3", args{[]int{1}, 0}, 0},
		{"Test 4", args{[]int{1}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange2(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange2() = %v, want %v", got, tt.want)
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

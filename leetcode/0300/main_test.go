package main

import (
	"testing"
)

func Test_lengthOfLIS1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
		{"case2", args{[]int{0, 1, 0, 3, 2, 3}}, 4},
		{"case3", args{[]int{7, 7, 7, 7, 7, 7, 7}}, 1},
		{"case4", args{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}}, 6},
		{"case5", args{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6, 8}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS1(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
		{"case2", args{[]int{0, 1, 0, 3, 2, 3}}, 4},
		{"case3", args{[]int{7, 7, 7, 7, 7, 7, 7}}, 1},
		{"case4", args{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}}, 6},
		{"case5", args{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6, 8}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS2(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS2() = %v, want %v", got, tt.want)
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
		{"case1", args{1, 2}, 2},
		{"case2", args{2, 1}, 2},
		{"case3", args{1, 1}, 1},
		{"case4", args{0, 0}, 0},
		{"case5", args{0, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
		{"case2", args{[]int{0, 1, 0, 3, 2, 3}}, 4},
		{"case3", args{[]int{7, 7, 7, 7, 7, 7, 7}}, 1},
		{"case4", args{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}}, 6},
		{"case5", args{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6, 8}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS3(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS3() = %v, want %v", got, tt.want)
			}
		})
	}
}

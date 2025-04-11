package main

import (
	"testing"
)

func Test_maxProfit1(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]int{1, 2, 3, 0, 2}}, 3},
		{"Case 2", args{[]int{1, 2, 4}}, 3},
		{"Case 3", args{[]int{1, 2, 3, 4, 5}}, 4},
		{"Case 4", args{[]int{1, 2, 3, 4, 5, 6}}, 5},
		{"Case 5", args{[]int{1, 2, 3, 4, 5, 6, 7}}, 6},
		{"Case 6", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}}, 7},
		{"Case 7", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit1(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit1() = %v, want %v", got, tt.want)
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
		{"Case 1", args{1, 2}, 2},
		{"Case 2", args{2, 1}, 2},
		{"Case 3", args{2, 2}, 2},
		{"Case 4", args{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]int{1, 2, 3, 0, 2}}, 3},
		{"Case 2", args{[]int{1, 2, 4}}, 3},
		{"Case 3", args{[]int{1, 2, 3, 4, 5}}, 4},
		{"Case 4", args{[]int{1, 2, 3, 4, 5, 6}}, 5},
		{"Case 5", args{[]int{1, 2, 3, 4, 5, 6, 7}}, 6},
		{"Case 6", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}}, 7},
		{"Case 7", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit2(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}

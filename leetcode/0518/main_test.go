package main

import "testing"

func Test_change1(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{5, []int{1, 2, 5}}, 4},
		{"Test 2", args{3, []int{2}}, 0},
		{"Test 3", args{10, []int{10}}, 1},
		{"Test 4", args{10, []int{1, 2, 5}}, 10},
		{"Test 5", args{500, []int{3, 5, 7, 8, 9, 10, 11}}, 35502874},
		{"Test 6", args{500, []int{1, 2, 5}}, 12701},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change1(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("change1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_change(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{5, []int{1, 2, 5}}, 4},
		{"Test 2", args{3, []int{2}}, 0},
		{"Test 3", args{10, []int{10}}, 1},
		{"Test 4", args{10, []int{1, 2, 5}}, 10},
		{"Test 5", args{500, []int{3, 5, 7, 8, 9, 10, 11}}, 35502874},
		{"Test 6", args{500, []int{1, 2, 5}}, 12701},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("change() = %v, want %v", got, tt.want)
			}
		})
	}
}

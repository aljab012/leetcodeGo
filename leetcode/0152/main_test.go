package main

import "testing"

func Test_maxProduct1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test case 1", args{[]int{2, 3, -2, 4}}, 6},
		{"Test case 2", args{[]int{-2, 0, -1}}, 0},
		{"Test case 3", args{[]int{-2, 3, -4}}, 24},
		{"Test case 4", args{[]int{2, -5, -2, -4, 3}}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct1(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProduct2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test case 1", args{[]int{2, 3, -2, 4}}, 6},
		{"Test case 2", args{[]int{-2, 0, -1}}, 0},
		{"Test case 3", args{[]int{-2, 3, -4}}, 24},
		{"Test case 4", args{[]int{2, -5, -2, -4, 3}}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct2(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test case 1", args{[]int{2, 3, -2, 4}}, 6},
		{"Test case 2", args{[]int{-2, 0, -1}}, 0},
		{"Test case 3", args{[]int{-2, 3, -4}}, 24},
		{"Test case 4", args{[]int{2, -5, -2, -4, 3}}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
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
		{"Test case 1", args{1, 2}, 2},
		{"Test case 2", args{2, 1}, 2},
		{"Test case 3", args{1, 1}, 1},
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
		{"Test case 1", args{1, 2}, 1},
		{"Test case 2", args{2, 1}, 1},
		{"Test case 3", args{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

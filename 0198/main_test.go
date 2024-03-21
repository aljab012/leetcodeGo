package main

import "testing"

func Test_rob1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 1}}, 4},
		{"2", args{[]int{2, 7, 9, 3, 1}}, 12},
		{"3", args{[]int{2, 1, 1, 2}}, 4},
		{"4", args{[]int{1, 2, 3, 1, 1, 100}}, 104},
		{"5", args{[]int{1, 2, 3, 1, 1, 100, 1}}, 104},
		{"6", args{[]int{1, 2, 3, 1, 1, 100, 1, 1}}, 105},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob1(tt.args.nums); got != tt.want {
				t.Errorf("rob1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rob2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 1}}, 4},
		{"2", args{[]int{2, 7, 9, 3, 1}}, 12},
		{"3", args{[]int{2, 1, 1, 2}}, 4},
		{"4", args{[]int{1, 2, 3, 1, 1, 100}}, 104},
		{"5", args{[]int{1, 2, 3, 1, 1, 100, 1}}, 104},
		{"6", args{[]int{1, 2, 3, 1, 1, 100, 1, 1}}, 105},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob2(tt.args.nums); got != tt.want {
				t.Errorf("rob2() = %v, want %v", got, tt.want)
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
		{"1", args{1, 2}, 2},
		{"2", args{2, 1}, 2},
		{"3", args{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_canPartition1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{[]int{1, 5, 11, 5}}, true},
		{"case2", args{[]int{1, 2, 3, 5}}, false},
		{"case3", args{[]int{1, 2, 5}}, false},
		{"case4", args{[]int{1, 2, 3, 4, 5, 6, 7}}, true},
		{"case5", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition1(tt.args.nums); got != tt.want {
				t.Errorf("canPartition1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPartition2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{[]int{1, 5, 11, 5}}, true},
		{"case2", args{[]int{1, 2, 3, 5}}, false},
		{"case3", args{[]int{1, 2, 5}}, false},
		{"case4", args{[]int{1, 2, 3, 4, 5, 6, 7}}, true},
		{"case5", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition2(tt.args.nums); got != tt.want {
				t.Errorf("canPartition2() = %v, want %v", got, tt.want)
			}
		})
	}
}

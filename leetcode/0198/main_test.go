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
		{"Example 1", args{[]int{1, 2, 3, 1}}, 4},
		{"Example 2", args{[]int{2, 7, 9, 3, 1}}, 12},
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

		{"Example 1", args{[]int{1, 2, 3, 1}}, 4},
		{"Example 2", args{[]int{2, 7, 9, 3, 1}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob2(tt.args.nums); got != tt.want {
				t.Errorf("rob2() = %v, want %v", got, tt.want)
			}
		})
	}
}

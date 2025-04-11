package main

import (
	"testing"
)

func Test_findTargetSumWays1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{nums: []int{1, 1, 1, 1, 1}, target: 3}, want: 5},
		{name: "Example 2", args: args{nums: []int{1}, target: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTargetSumWays2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{nums: []int{1, 1, 1, 1, 1}, target: 3}, want: 5},
		{name: "Example 2", args: args{nums: []int{1}, target: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays2() = %v, want %v", got, tt.want)
			}
		})
	}
}

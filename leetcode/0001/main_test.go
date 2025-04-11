package main

import (
	"reflect"
	"testing"
)

func Test_twoSum1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test case 1", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"test case 2", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"test case 3", args{[]int{3, 3}, 6}, []int{0, 1}},
		{"test case 4", args{[]int{3, 2, 3}, 6}, []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum1(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test case 1", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"test case 2", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"test case 3", args{[]int{3, 3}, 6}, []int{0, 1}},
		{"test case 4", args{[]int{3, 2, 3}, 6}, []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum2(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum3(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test case 1", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"test case 2", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"test case 3", args{[]int{3, 3}, 6}, []int{0, 1}},
		{"test case 4", args{[]int{3, 2, 3}, 6}, []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum3(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}

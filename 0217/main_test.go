package main

import (
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Contains Duplicate",
			args: args{nums: []int{1, 2, 3, 1}},
			want: true,
		},
		{
			name: "No Duplicate",
			args: args{nums: []int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "Contains Duplicate",
			args: args{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate1(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsDuplicate1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Contains Duplicate",
			args: args{nums: []int{1, 2, 3, 1}},
			want: true,
		},
		{
			name: "No Duplicate",
			args: args{nums: []int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "Contains Duplicate",
			args: args{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate1(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsDuplicate2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate2(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate2() = %v, want %v", got, tt.want)
			}
		})
	}
}

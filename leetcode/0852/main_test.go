package main

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{arr: []int{0, 1, 0}},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{arr: []int{0, 2, 1, 0}},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{arr: []int{0, 10, 5, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

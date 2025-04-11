package main

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		h     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				piles: []int{3, 6, 7, 11},
				h:     8,
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				piles: []int{30, 11, 23, 4, 20},
				h:     5,
			},
			want: 30,
		},
		{
			name: "Example 3",
			args: args{
				piles: []int{30, 11, 23, 4, 20},
				h:     6,
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.h); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

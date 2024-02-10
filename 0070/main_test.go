package main

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 0 stairs", args{1}, 1},
		{"Test 1 stairs", args{2}, 2},
		{"Test 2 stairs ", args{3}, 3},
		{"Test 3 stairs", args{4}, 5},
		{"Test 4 stairs", args{5}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

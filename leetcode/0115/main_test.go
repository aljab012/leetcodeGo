package main

import "testing"

func Test_numDistinct1(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{"rabbbit", "rabbit"},
			3,
		},
		{
			"Example 2",
			args{"babgbag", "bag"},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct1(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinct1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numDistinct2(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{"rabbbit", "rabbit"},
			3,
		},
		{
			"Example 2",
			args{"babgbag", "bag"},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinct2() = %v, want %v", got, tt.want)
			}
		})
	}
}

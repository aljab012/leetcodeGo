package main

import "testing"

func Test_isInterleave1(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Example 1",
			args{"aabcc", "dbbca", "aadbbcbcac"},
			true,
		},
		{
			"Example 2",
			args{"aabcc", "dbbca", "aadbbbaccc"},
			false,
		},
		{
			"Example 3",
			args{"", "", ""},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave1(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isInterleave2(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Example 1",
			args{"aabcc", "dbbca", "aadbbcbcac"},
			true,
		},
		{
			"Example 2",
			args{"aabcc", "dbbca", "aadbbbaccc"},
			false,
		},
		{
			"Example 3",
			args{"", "", ""},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave2(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave2() = %v, want %v", got, tt.want)
			}
		})
	}
}

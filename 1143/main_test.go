package main

import "testing"

func Test_longestCommonSubsequence1(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{"abcde", "ace"}, 3},
		{"Test 2", args{"abc", "abc"}, 3},
		{"Test 3", args{"abc", "def"}, 0},
		{"Test 4", args{"abc", ""}, 0},
		{"Test 5", args{"", "abc"}, 0},
		{"Test 6", args{"", ""}, 0},
		{"Test 7", args{"bsbininm", "jmjkbkjkv"}, 1},
		{"Test 8", args{"abcba", "abcbcba"}, 5},
		{"Test 9", args{"abcba", "abcbcba"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence1(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubsequence2(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{"abcde", "ace"}, 3},
		{"Test 2", args{"abc", "abc"}, 3},
		{"Test 3", args{"abc", "def"}, 0},
		{"Test 4", args{"abc", ""}, 0},
		{"Test 5", args{"", "abc"}, 0},
		{"Test 6", args{"", ""}, 0},
		{"Test 7", args{"bsbininm", "jmjkbkjkv"}, 1},
		{"Test 8", args{"abcba", "abcbcba"}, 5},
		{"Test 9", args{"abcba", "abcbcba"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence2(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{1, 2}, 2},
		{"Test 2", args{2, 1}, 2},
		{"Test 3", args{1, 1}, 1},
		{"Test 4", args{0, 0}, 0},
		{"Test 5", args{0, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

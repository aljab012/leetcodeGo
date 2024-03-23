package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{"babad"}, "bab"},
		{"Example 2", args{"cbbd"}, "bb"},
		{"Example 3", args{"a"}, "a"},
		{"Example 4", args{"ac"}, "a"},
		{"Example 5", args{"bb"}, "bb"},
		{"Example 6", args{"ccc"}, "ccc"},
		{"Example 7", args{"aaaa"}, "aaaa"},
		{"Example 8", args{"aaaaa"}, "aaaaa"},
		{"Example 9", args{"aaaaaa"}, "aaaaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

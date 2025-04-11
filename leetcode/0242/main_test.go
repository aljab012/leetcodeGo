package main

import (
	"testing"
)

func Test_isAnagram1(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{"anagram", "nagaram"}, true,},
		{"case2", args{"rat", "car"}, false,},
		{"case3", args{"a", "ab"}, false,},
		{"case4", args{"a", "a"}, true,},
		{"case5", args{"a", "b"}, false,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram1(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagram2(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{"anagram", "nagaram"}, true,},
		{"case2", args{"rat", "car"}, false,},
		{"case3", args{"a", "ab"}, false,},
		{"case4", args{"a", "a"}, true,},
		{"case5", args{"a", "b"}, false,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortString(tt.args.s); got != tt.want {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}

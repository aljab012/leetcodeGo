package main

import (
	"testing"
)

func Test_numDecodings1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{"12"}, 2},
		{"Test 2", args{"226"}, 3},
		{"Test 3", args{"0"}, 0},
		{"Test 4", args{"06"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings1(tt.args.s); got != tt.want {
				t.Errorf("numDecodings1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numDecodings2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{"12"}, 2},
		{"Test 2", args{"226"}, 3},
		{"Test 3", args{"0"}, 0},
		{"Test 4", args{"06"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings2(tt.args.s); got != tt.want {
				t.Errorf("numDecodings2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numDecodings3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{"12"}, 2},
		{"Test 2", args{"226"}, 3},
		{"Test 3", args{"0"}, 0},
		{"Test 4", args{"06"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings3(tt.args.s); got != tt.want {
				t.Errorf("numDecodings3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numDecodings4(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{"12"}, 2},
		{"Test 2", args{"226"}, 3},
		{"Test 3", args{"0"}, 0},
		{"Test 4", args{"06"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings4(tt.args.s); got != tt.want {
				t.Errorf("numDecodings4() = %v, want %v", got, tt.want)
			}
		})
	}
}

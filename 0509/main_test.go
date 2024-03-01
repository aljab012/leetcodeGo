package main

import (
	"testing"
)

func Test_fib1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{2}, 1},
		{"Test 2", args{3}, 2},
		{"Test 3", args{4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib1(tt.args.n); got != tt.want {
				t.Errorf("fib1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{2}, 1},
		{"Test 2", args{3}, 2},
		{"Test 3", args{4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib2(tt.args.n); got != tt.want {
				t.Errorf("fib2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{2}, 1},
		{"Test 2", args{3}, 2},
		{"Test 3", args{4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib3(tt.args.n); got != tt.want {
				t.Errorf("fib3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib4(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{2}, 1},
		{"Test 2", args{3}, 2},
		{"Test 3", args{4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib4(tt.args.n); got != tt.want {
				t.Errorf("fib4() = %v, want %v", got, tt.want)
			}
		})
	}
}

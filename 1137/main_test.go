package main

import "testing"

func Test_tribonacci1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{4}, 4},
		{"Test 2", args{25}, 1389537},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci1(tt.args.n); got != tt.want {
				t.Errorf("tribonacci1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tribonacci2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{4}, 4},
		{"Test 2", args{25}, 1389537},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci2(tt.args.n); got != tt.want {
				t.Errorf("tribonacci2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tribonacci3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{4}, 4},
		{"Test 2", args{25}, 1389537},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci3(tt.args.n); got != tt.want {
				t.Errorf("tribonacci3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tribonacci4(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{4}, 4},
		{"Test 2", args{25}, 1389537},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci4(tt.args.n); got != tt.want {
				t.Errorf("tribonacci4() = %v, want %v", got, tt.want)
			}
		})
	}
}

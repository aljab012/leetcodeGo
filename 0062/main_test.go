package main

import "testing"

func Test_uniquePaths1(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{3, 7}, 28},
		{"Test 2", args{3, 2}, 3},
		{"Test 3", args{7, 3}, 28},
		{"Test 4", args{3, 3}, 6},
		{"Test 5", args{3, 4}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths1(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePaths2(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{3, 7}, 28},
		{"Test 2", args{3, 2}, 3},
		{"Test 3", args{7, 3}, 28},
		{"Test 4", args{3, 3}, 6},
		{"Test 5", args{3, 4}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths2(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePaths3(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{3, 7}, 28},
		{"Test 2", args{3, 2}, 3},
		{"Test 3", args{7, 3}, 28},
		{"Test 4", args{3, 3}, 6},
		{"Test 5", args{3, 4}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths3(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths3() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_uniquePaths4(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{3, 7}, 28},
		{"Test 2", args{3, 2}, 3},
		{"Test 3", args{7, 3}, 28},
		{"Test 4", args{3, 3}, 6},
		{"Test 5", args{3, 4}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths3(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths3() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_canVisitAllRooms1(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Example 1",
			args{[][]int{{1}, {2}, {3}, {}}},
			true,
		},
		{
			"Example 2",
			args{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canVisitAllRooms1(tt.args.rooms); got != tt.want {
				t.Errorf("canVisitAllRooms1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canVisitAllRooms2(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Example 1",
			args{[][]int{{1}, {2}, {3}, {}}},
			true,
		},
		{
			"Example 2",
			args{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canVisitAllRooms2(tt.args.rooms); got != tt.want {
				t.Errorf("canVisitAllRooms2() = %v, want %v", got, tt.want)
			}
		})
	}
}

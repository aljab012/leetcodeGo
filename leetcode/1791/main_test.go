package main

import (
	"testing"
)

func Test_findCenter(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{edges: [][]int{{1, 2}, {2, 3}, {4, 2}}},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{edges: [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCenter(tt.args.edges); got != tt.want {
				t.Errorf("findCenter() = %v, want %v", got, tt.want)
			}
		})
	}
}

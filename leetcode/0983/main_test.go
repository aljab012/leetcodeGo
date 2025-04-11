package main

import "testing"

func Test_mincostTickets(t *testing.T) {
	type args struct {
		days  []int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				days:  []int{1, 4, 6, 7, 8, 20},
				costs: []int{2, 7, 15},
			},
			want: 11,
		},
		{
			name: "Example 2",
			args: args{
				days:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
				costs: []int{2, 7, 15},
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostTickets1(tt.args.days, tt.args.costs); got != tt.want {
				t.Errorf("mincostTickets1() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostTickets2(tt.args.days, tt.args.costs); got != tt.want {
				t.Errorf("mincostTickets2() = %v, want %v", got, tt.want)
			}
		})
	}
}

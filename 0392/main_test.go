package main

import "testing"

func Test_isSubsequence1(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence1(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence1() = %v, want %v", got, tt.want)
			}
		})
	}
}

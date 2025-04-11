package main

import (
	"reflect"
	"testing"
)

func Test_wordBreak1(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			want: []string{"leet code"},
		},
		{
			name: "Example 2",
			args: args{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			want: []string{"apple pen apple"},
		},
		{
			name: "Example 3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak1(tt.args.s, tt.args.wordDict); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordBreak1() = %v, want %v", got, tt.want)
			}
		})
	}
}

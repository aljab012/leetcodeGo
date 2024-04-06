package main

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"Test case 1", args{[]string{"a"}}, [][]string{{"a"}}},
		{"Test case 2", args{[]string{"a", "b"}}, [][]string{{"a"}, {"b"}}},
		{"Test case 3", args{[]string{"a", "b", "c"}}, [][]string{{"a"}, {"b"}, {"c"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

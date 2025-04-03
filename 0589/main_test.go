package main

import (
	"reflect"
	"testing"
)

func Test_preorder1(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorder1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorder2(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorder2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_sumRootToLeaf(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumRootToLeaf(tt.args.root); got != tt.want {
				t.Errorf("sumRootToLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}

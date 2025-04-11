package main

import "testing"

func Test_hasCycle(t *testing.T) {

	// Example 1
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2

	// Example 2
	c1 := &ListNode{Val: 1}
	c2 := &ListNode{Val: 2}
	c1.Next = c2
	c2.Next = c1

	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{"Multiple nodes with cycle", n1, true},
		{"Two nodes with cycle", c1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle1(tt.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle1(tt.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

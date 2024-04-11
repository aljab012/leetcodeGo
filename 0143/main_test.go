package main

import "testing"

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name     string
		args     args
		expected *ListNode
	}{
		{
			name: "Test 1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "Test 2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "Test 3",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
		{
			name: "Test 4",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			expected: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
		{
			name: "Test 5",
			args: args{
				head: nil,
			},
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			for tt.expected != nil {
				if tt.args.head.Val != tt.expected.Val {
					t.Errorf("reorderList() = %v, want %v", tt.args.head.Val, tt.expected.Val)
				}
				tt.args.head = tt.args.head.Next
				tt.expected = tt.expected.Next
			}
		})
	}
}

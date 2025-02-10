package main

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "test case 1",
			args: args{
				head: &Node{
					Val: 7,
					Next: &Node{
						Val: 13,
						Next: &Node{
							Val: 11,
							Next: &Node{
								Val: 10,
								Next: &Node{
									Val: 1,
								},
							},
						},
					},
				},
			},
			want: &Node{
				Val: 7,
				Next: &Node{
					Val: 13,
					Next: &Node{
						Val: 11,
						Next: &Node{
							Val: 10,
							Next: &Node{
								Val: 1,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}

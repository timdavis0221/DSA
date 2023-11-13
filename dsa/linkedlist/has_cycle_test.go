package linkedlist

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "HasCycle_Test1",
			args: args{
				head: SinglyLinkedList([]int{1, 2, 3}),
			},
			want: false,
		}, {
			name: "HasCycle_Test2",
			args: args{
				head: CycledLinkedList([]int{1, 2, 3}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycle(tt.args.head); got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasCycleApproach2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "HasCycleApproach2_Test1",
			args: args{
				head: SinglyLinkedList([]int{1, 2, 3}),
			},
			want: false,
		}, {
			name: "HasCycleApproach2_Test2",
			args: args{
				head: CycledLinkedList([]int{1, 2, 3}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycleApproach2(tt.args.head); got != tt.want {
				t.Errorf("HasCycleApproach2() = %v, want %v", got, tt.want)
			}
		})
	}
}

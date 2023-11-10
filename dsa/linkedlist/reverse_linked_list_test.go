package linkedlist

import (
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ReverseLinkedList_Test1",
			args: args{
				head: CreateLinkedList([]int{1, 2, 3, 4, 5}),
			},
			want: CreateLinkedList([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseLinkedList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseLinkedListRecursively(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ReverseLinkedListRecursively_Test1",
			args: args{
				head: CreateLinkedList([]int{1, 2, 3, 4, 5}),
			},
			want: CreateLinkedList([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseLinkedListRecursively(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseLinkedListRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}

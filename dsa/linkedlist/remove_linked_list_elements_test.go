package linkedlist

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "RemoveElements_Test1",
			args: args{
				head: CreateLinkedList([]int{1, 2, 6, 3, 4, 5, 6}),
				val:  6,
			},
			want: CreateLinkedList([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

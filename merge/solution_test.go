package merge

import (
	"reflect"
	"testing"
)

func CreateList(nums *[]int) *ListNode {
	arr := *nums
	var head *ListNode
	var cur *ListNode
	for idx, val := range arr {
		if idx == 0 {
			head = &ListNode{Val: val}
			cur = head
		} else {
			cur.Next = &ListNode{Val: val}
			cur = cur.Next
		}
	}
	return head
}
func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "list1 = [1,2,4], list2 = [1,3,4]",
			args: args{l1: CreateList(&[]int{1, 2, 4}), l2: CreateList(&[]int{1, 3, 4})},
			want: CreateList(&[]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "list1 = [], list2 = []",
			args: args{l1: CreateList(&[]int{}), l2: CreateList(&[]int{})},
			want: CreateList(&[]int{}),
		},
		{
			name: "list1 = [], list2 = [0]",
			args: args{l1: CreateList(&[]int{}), l2: CreateList(&[]int{0})},
			want: CreateList(&[]int{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

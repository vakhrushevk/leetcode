package mergeTwoSortedList

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func NewNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func NodeTest1() *ListNode {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(5)
	node6 := NewNode(6)
	node7 := NewNode(7)
	node1 = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node3
	return node1
}

func NodeTest(nums ...int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, i := range nums {
		cur.Next = &ListNode{Val: i}
		cur = cur.Next
	}
	return head.Next
}

func toArray(head *ListNode) []int {
	current := head
	result := []int{}
	for current != nil {
		result = append(result, current.Val)
		fmt.Println(result)
		current = current.Next
	}

	return result
}

func Test_LinkedListCycle(t *testing.T) {
	tests := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  []int
	}{
		{name: "Test case 0",
			list1: NodeTest(1, 2, 4),
			list2: NodeTest(1, 3, 4),
			want:  []int{1, 1, 2, 3, 4, 4}},
		{
			name:  "Test case 1",
			list1: NodeTest(1, 1, 1, 1),
			list2: NodeTest(1, 1, 1, 1),
			want:  []int{1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name:  "Test case 2",
			list1: NodeTest(1, 2, 3, 4),
			list2: NodeTest(5),
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Test case 3",
			list1: NodeTest(1),
			list2: NodeTest(2, 3, 4, 5),
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Test case 4",
			list1: nil,
			list2: NodeTest(5, 6, 7, 8),
			want:  []int{5, 6, 7, 8},
		},
		{
			name:  "Test case 5",
			list1: NodeTest(1, 2, 3, 4),
			list2: nil,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "Test case 6",
			list1: nil,
			list2: nil,
			want:  []int{},
		},
		{
			name:  "Test case 7",
			list1: NodeTest(1, 2, 3, 4, 5, 6, 9, 10),
			list2: NodeTest(5, 6, 7, 8),
			want:  []int{1, 2, 3, 4, 5, 5, 6, 6, 7, 8, 9, 10},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println("list1")
			toArray(tc.list1)
			fmt.Println("list2")
			toArray(tc.list2)
			result := mergeTwoLists(tc.list1, tc.list2)
			resArrey := toArray(result)
			require.Equal(t, resArrey, tc.want)
		})
	}
}

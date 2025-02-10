package LinkedListCycle

import (
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

func NodeTest2() *ListNode {
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
	node7.Next = nil
	return node1
}

func Test_LinkedListCycle(t *testing.T) {
	tests := []struct {
		name     string
		list     *ListNode
		expected bool
	}{
		{
			name:     "Test case 0",
			list:     nil,
			expected: false,
		},
		{
			name:     "Test case 1",
			list:     NodeTest1(),
			expected: true,
		},
		{
			name:     "Test case 2",
			list:     NodeTest2(),
			expected: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := hasCycle(tc.list)
			require.Equal(t, tc.expected, result)
		})
	}
}

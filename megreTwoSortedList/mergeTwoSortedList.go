package mergeTwoSortedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	damy := &ListNode{}
	result := damy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			damy.Next = list1
			list1 = list1.Next
		} else {
			damy.Next = list2
			list2 = list2.Next
		}
		damy = damy.Next
	}
	if list1 != nil {
		damy.Next = list1
	}
	if list2 != nil {
		damy.Next = list2
	}

	return result.Next
}

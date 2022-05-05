package merge

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var merged *ListNode
	var cur *ListNode
	if l1 == nil && l2 == nil {
		return merged
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	for l1 != nil && l2 != nil {
		if merged == nil {
			if l1.Val < l2.Val {
				merged = &ListNode{Val: l1.Val}
				l1 = l1.Next
			} else {
				merged = &ListNode{Val: l2.Val}
				l2 = l2.Next
			}
			cur = merged
		} else {
			if l1.Val < l2.Val {
				cur.Next = &ListNode{Val: l1.Val}
				l1 = l1.Next
			} else {
				cur.Next = &ListNode{Val: l2.Val}
				l2 = l2.Next
			}
			cur = cur.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return merged
}

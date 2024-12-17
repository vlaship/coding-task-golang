package problem_0021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	cur := &ListNode{}
	res := cur

	for {
		l, r := 0, 0

		if list1 != nil {
			l = list1.Val
		}
		if list2 != nil {
			r = list2.Val
		}

		if l < r && list1 != nil {
			cur.Val = l
			list1 = list1.Next
		} else if list2 != nil {
			cur.Val = r
			list2 = list2.Next
		} else {
			cur.Val = l
			list1 = list1.Next
		}

		if list2 != nil || list1 != nil {
			cur.Next = &ListNode{}
			cur = cur.Next
		} else {
			return res
		}
	}
}

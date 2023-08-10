package problem0002

/*
2. Add Two Numbers

https://leetcode.com/problems/add-two-numbers/

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:
The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r1 := getNum(l1.Next, uint64(l1.Val))
	r2 := getNum(l2.Next, uint64(l2.Val))

	return getListNode(r1 + r2)
}

func getNum(l *ListNode, num uint64) uint64 {
	if l == nil {
		return num
	}
	i := getNum(l.Next, uint64(l.Val))*uint64(10) + num
	return i
}

func getListNode(num uint64) *ListNode {
	if num < 10 {
		return &ListNode{Val: int(num)}
	}
	return &ListNode{Val: int(num % 10), Next: getListNode(num / 10)}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

package problem_0002

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addTwoNumbers_1(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	expected := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}

	result := addTwoNumbers(l1, l2)
	require.Equal(t, expected, result)
}

func Test_addTwoNumbers_2(t *testing.T) {
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 0}
	expected := &ListNode{Val: 0}

	result := addTwoNumbers(l1, l2)
	require.Equal(t, expected, result)
}

func Test_addTwoNumbers_3(t *testing.T) {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9,
		Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9,
			Next: &ListNode{Val: 9}}}}}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9,
		Next: &ListNode{Val: 9}}}}
	expected := &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9,
		Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
			Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}}}

	result := addTwoNumbers(l1, l2)
	require.Equal(t, expected, result)
}

func Test_addTwoNumbers_4(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}}
	expected := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}

	result := addTwoNumbers(l1, l2)
	require.Equal(t, expected, result)
}

func Test_addTwoNumbers_5(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
		Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
			Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
				Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
					Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
						Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
							Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
								Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
									Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
										Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
											Next: &ListNode{Val: 1}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}

	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	expected := &ListNode{Val: 6, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4,
		Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
			Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
				Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
					Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
						Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
							Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
								Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
									Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
										Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0,
											Next: &ListNode{Val: 1}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}
	result := addTwoNumbers(l1, l2)
	require.Equal(t, expected, result)
}

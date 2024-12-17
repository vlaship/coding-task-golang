package problem_0021

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Helper function to create a linked list from a slice
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

// Helper function to convert linked list to slice for comparison
func listToSlice(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{
			name:  "Example 1",
			list1: []int{1, 2, 4},
			list2: []int{1, 3, 4},
			want:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:  "Example 2",
			list1: []int{},
			list2: []int{},
			want:  []int{},
		},
		{
			name:  "Example 3",
			list1: []int{},
			list2: []int{0},
			want:  []int{0},
		},
		{
			name:  "One empty list",
			list1: []int{1, 2, 3},
			list2: []int{},
			want:  []int{1, 2, 3},
		},
		{
			name:  "Different lengths",
			list1: []int{1, 2},
			list2: []int{3, 4, 5, 6},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := createList(tt.list1)
			l2 := createList(tt.list2)
			got := mergeTwoLists(l1, l2)
			gotSlice := listToSlice(got)
			if len(tt.want) == 0 && len(gotSlice) == 0 {
				return // both empty, pass
			}
			require.Equal(t, tt.want, gotSlice)
		})
	}
}

package problem_0707

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyLinkedList(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		myLinkedList := Constructor()
		myLinkedList.AddAtHead(1)
		myLinkedList.AddAtTail(3)
		myLinkedList.AddAtIndex(1, 2) // linked list becomes 1->2->3

		assert.Equal(t, 2, myLinkedList.Get(1)) // return 2

		myLinkedList.DeleteAtIndex(1) // now the linked list is 1->3

		assert.Equal(t, 3, myLinkedList.Get(1)) // return 3
	})
}

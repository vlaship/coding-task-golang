package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTask3(t *testing.T) {
	expected := [][]int{{3, 3}, {1, 5}, {4, 2}}
	result := Task3([]int{1, 3, 3, 4, 5, 2}, 6)
	if !assert.Equal(t, result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

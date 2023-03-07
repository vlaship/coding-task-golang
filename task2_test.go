package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTask2(t *testing.T) {
	expected := []int{5, 6}
	result := Task2([]int{4, 3, 2, 7, 8, 2, 3, 1})
	if !assert.Equal(t, result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

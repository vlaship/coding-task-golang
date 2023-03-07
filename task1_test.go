package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTask1(t *testing.T) {
	expected := []int{0, 1, 9, 16, 16, 100}
	result := Task1([]int{-4, -1, 0, 3, 4, 10})
	if !assert.Equal(t, result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

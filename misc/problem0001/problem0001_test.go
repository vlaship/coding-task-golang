package problem0001

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_problem0001(t *testing.T) {
	expected := []int{0, 1, 9, 16, 16, 100}
	result := problem001([]int{-4, -1, 0, 3, 4, 10})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

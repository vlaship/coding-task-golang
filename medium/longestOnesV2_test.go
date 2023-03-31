package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestOnesV2_1(t *testing.T) {
	expected := 6
	result := longestOnesV2([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}
func Test_longestOnesV2_2(t *testing.T) {
	expected := 10
	result := longestOnesV2([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

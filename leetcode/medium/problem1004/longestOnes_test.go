package problem1004

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestOnes_1(t *testing.T) {
	expected := 6
	result := longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}
func Test_longestOnes_2(t *testing.T) {
	expected := 10
	result := longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

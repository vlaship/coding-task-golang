package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum_1(t *testing.T) {
	expected := []int{0, 1}
	result := twoSum([]int{2, 7, 11, 15}, 9)
	if !assert.Equal(t, result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func Test_twoSum_2(t *testing.T) {
	expected := []int{1, 2}
	result := twoSum([]int{3, 2, 4}, 6)
	if !assert.Equal(t, result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func Test_twoSum_3(t *testing.T) {
	expected := []int{0, 1}
	result := twoSum([]int{3, 3}, 6)
	if !assert.Equal(t, result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

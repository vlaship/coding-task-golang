package problem0238

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productExceptSelf_1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}
	result := productExceptSelf(nums)

	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

func Test_productExceptSelf_2(t *testing.T) {
	nums := []int{-1, 1, 0, -3, 3}
	expected := []int{0, 0, 9, 0, 0}
	result := productExceptSelf(nums)

	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findDisappearedNumbers_1(t *testing.T) {
	expected := []int{5, 6}
	result := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func Test_findDisappearedNumbers_2(t *testing.T) {
	expected := []int{2}
	result := findDisappearedNumbers([]int{1, 1})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

package problem0217

import (
	"testing"
)

func Test_containsDuplicate_1(t *testing.T) {
	result := containsDuplicate([]int{1, 2, 3, 1})
	if !result {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

func Test_containsDuplicate_2(t *testing.T) {
	result := containsDuplicate([]int{1, 2, 3, 4})
	if result {
		t.Errorf("Expected %v but got %v", false, result)
	}
}

func Test_containsDuplicate_3(t *testing.T) {
	result := containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	if !result {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

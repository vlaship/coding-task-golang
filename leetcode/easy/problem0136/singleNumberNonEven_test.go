package problem0136

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumberNonOnlyEven_1(t *testing.T) {
	expected := 1
	result := singleNumberNonOnlyEven([]int{2, 2, 1})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func Test_singleNumberNonOnlyEven_2(t *testing.T) {
	expected := 4
	result := singleNumberNonOnlyEven([]int{4, 1, 2, 1, 2})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func Test_singleNumberNonOnlyEven_3(t *testing.T) {
	expected := 1
	result := singleNumberNonOnlyEven([]int{1})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func Test_singleNumberNonOnlyEven_4(t *testing.T) {
	expected := 1
	result := singleNumberNonOnlyEven([]int{2, 2, 1, 2})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

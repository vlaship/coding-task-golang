package problem0268

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_missingNumber_1(t *testing.T) {
	expected := 2
	result := missingNumber([]int{3, 0, 1})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func Test_missingNumber_2(t *testing.T) {
	expected := 2
	result := missingNumber([]int{0, 1})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func Test_missingNumber_3(t *testing.T) {
	expected := 8
	result := missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

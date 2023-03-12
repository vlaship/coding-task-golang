package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber_1(t *testing.T) {
	expected := 1
	result := singleNumber([]int{2, 2, 1})
	if !assert.Equal(t, result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func Test_singleNumber_2(t *testing.T) {
	expected := 4
	result := singleNumber([]int{4, 1, 2, 1, 2})
	if !assert.Equal(t, result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func Test_singleNumber_3(t *testing.T) {
	expected := 1
	result := singleNumber([]int{1})
	if !assert.Equal(t, result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

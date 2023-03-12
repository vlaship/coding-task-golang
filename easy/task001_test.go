package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_task001(t *testing.T) {
	expected := []int{0, 1, 9, 16, 16, 100}
	result := task001([]int{-4, -1, 0, 3, 4, 10})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

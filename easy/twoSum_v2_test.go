package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSumV2(t *testing.T) {
	expected := [][]int{{3, 3}, {1, 5}, {4, 2}}
	result := twoSumV2([]int{1, 3, 3, 4, 5, 2}, 6)
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

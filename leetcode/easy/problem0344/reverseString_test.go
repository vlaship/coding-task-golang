package problem0344

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseString(t *testing.T) {
	expected := "!dlroW olleH"
	result := reverseString("Hello World!")
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

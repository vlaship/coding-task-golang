package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTask4(t *testing.T) {
	expected := "!dlroW olleH"
	result := Task4("Hello World!")
	if !assert.Equal(t, result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

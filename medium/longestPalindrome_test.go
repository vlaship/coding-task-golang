package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome_1(t *testing.T) {
	expected := "bab"
	result := longestPalindrome("babad")
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

func Test_longestPalindrome_2(t *testing.T) {
	expected := "bb"
	result := longestPalindrome("cbbd")
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

func Test_longestPalindrome_3(t *testing.T) {
	expected := "aaaa"
	result := longestPalindrome("aaaa")
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

func Test_longestPalindrome_4(t *testing.T) {
	expected := "aca"
	result := longestPalindrome("aacabdkacaa")
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

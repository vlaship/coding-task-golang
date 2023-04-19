package easy

import (
	"testing"
)

func Test_isValidParentheses_1(t *testing.T) {
	result := isValidParentheses("()")
	if !result {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

func Test_isValidParentheses_2(t *testing.T) {
	result := isValidParentheses("()[]{}")
	if !result {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

func Test_isValidParentheses_3(t *testing.T) {
	result := isValidParentheses("(]")
	if result {
		t.Errorf("Expected %v but got %v", false, result)
	}
}

func Test_isValidParentheses_4(t *testing.T) {
	result := isValidParentheses("){")
	if result {
		t.Errorf("Expected %v but got %v", false, result)
	}
}

func Test_isValidParentheses_5(t *testing.T) {
	result := isValidParentheses("(){}}{")
	if result {
		t.Errorf("Expected %v but got %v", false, result)
	}
}

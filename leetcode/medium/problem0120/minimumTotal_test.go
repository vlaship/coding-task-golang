package problem0120

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumTotal_1(t *testing.T) {
	expected := 11
	result := minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}
func Test_minimumTotal_2(t *testing.T) {
	expected := -10
	result := minimumTotal([][]int{{-10}})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

//2(2)
//3(5) 4(6)
//6(11) 5(10,11->10) 9(15)
//4(15) 4(15,14->14) 8(19,23->19) 0(15)
//5(20) 6(21,21->21) 7(21,26->21) 8(31,23->23) 9(24)

func Test_minimumTotal_3(t *testing.T) {
	expected := 14
	result := minimumTotal([][]int{{2}, {3, 4}, {6, 5, 9}, {4, 4, 8, 0}})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}
func Test_minimumTotal_4(t *testing.T) {
	expected := 20
	result := minimumTotal([][]int{{2}, {3, 4}, {6, 5, 9}, {4, 4, 8, 0}, {5, 6, 7, 8, 9}})
	if !assert.Equal(t, expected, result) {
		t.Errorf("Expected %v but got %v", true, result)
	}
}

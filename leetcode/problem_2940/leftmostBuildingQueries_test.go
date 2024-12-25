package problem_2940

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func Test_leftmostBuildingQueries(t *testing.T) {
    tests := []struct {
        name    string
        heights []int
        queries [][]int
        want    []int
    }{
        {
            name:    "example 1",
            heights: []int{6, 4, 8, 5, 2, 7},
            queries: [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}},
            want:    []int{2, 5, -1, 5, 2},
        },
        {
            name:    "example 2",
            heights: []int{5, 3, 8, 2, 6, 1, 4, 6},
            queries: [][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}},
            want:    []int{7, 6, -1, 4, 6},
        },
        {
            name: "large_test",
            heights: func() []int {
                h := make([]int, 10000)
                h[0] = 50000
                for i := 1; i < 10000; i++ {
                    h[i] = i
                }
                return h
            }(),
            queries: [][]int{
                {0, 1}, {0, 9999}, {9998, 9999}, {1, 2}, {1, 9999},
            },
            want:    []int{-1, -1, 9999, 2, 9999},
        },
        {
            name: "large_generated_test",
            heights: func() []int {
                h := make([]int, 10000)
                // First building very tall
                h[0] = 50000
                // Rest increasing linearly
                for i := 1; i < 10000; i++ {
                    h[i] = i
                }
                return h
            }(),
            queries: [][]int{
                {0, 1},     // Can't find taller than h[0]=50000
                {0, 9999},  // Can't find taller than h[0]=50000
                {9998, 9999}, // Can move to 9999 directly
                {1, 2},     // Can move to 2 directly
                {1, 9999},  // Can move to 9999 directly
            },
            want: []int{-1, -1, 9999, 2, 9999},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := leftmostBuildingQueries(tt.heights, tt.queries)
            require.Equal(t, tt.want, got)
        })
    }
}

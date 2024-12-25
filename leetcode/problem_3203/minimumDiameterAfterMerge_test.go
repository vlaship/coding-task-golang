package minimum_diameter_after_merge

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func Test_minimumDiameterAfterMerge(t *testing.T) {
    tests := []struct {
        name   string
        edges1 [][]int
        edges2 [][]int
        want   int
    }{
        {
            name:   "example 1",
            edges1: [][]int{{0, 1}, {0, 2}, {0, 3}},
            edges2: [][]int{{0, 1}},
            want:   3,
        },
        {
            name:   "example 2",
            edges1: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}},
            edges2: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}},
            want:   5,
        },
        {
            name:   "example 3 - failing case",
            edges1: [][]int{{0,1}, {2,0}, {3,2}, {3,6}, {8,7}, {4,8}, {5,4}, {3,5}, {3,9}},
            edges2: [][]int{{0,1}, {0,2}, {0,3}},
            want:   7,
        },
        {
            name:   "single nodes",
            edges1: [][]int{},
            edges2: [][]int{},
            want:   1,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := minimumDiameterAfterMerge(tt.edges1, tt.edges2)
            require.Equal(t, tt.want, got)
        })
    }
}

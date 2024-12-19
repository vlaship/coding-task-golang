package problem_0769

import "testing"

func Test_maxChunksToSorted(t *testing.T) {
    tests := []struct {
        name string
        arr  []int
        want int
    }{
        {
            name: "Example 1",
            arr:  []int{4, 3, 2, 1, 0},
            want: 1,
        },
        {
            name: "Example 2",
            arr:  []int{1, 0, 2, 3, 4},
            want: 4,
        },
        {
            name: "Single element",
            arr:  []int{0},
            want: 1,
        },
        {
            name: "Already sorted",
            arr:  []int{0, 1, 2, 3},
            want: 4,
        },
        {
            name: "Two chunks possible",
            arr:  []int{1, 0, 3, 2},
            want: 2,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := maxChunksToSorted(tt.arr); got != tt.want {
                t.Errorf("maxChunksToSorted() = %v, want %v", got, tt.want)
            }
        })
    }
}

package problem_0253

import (
	"container/heap"
	"sort"
)

type MinHeap []int

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Len() int           { return len(h) }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }

func (h *MinHeap) Pop() any {
	old := *h
	oldLen := len(old)
	last := old[oldLen-1]
	*h = old[0 : oldLen-1]
	return last
}
func minMeetingRooms(intervals [][]int) int {
	// sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// init
	h := &MinHeap{}
	heap.Init(h)

	// add first meeting
	h.Push(intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		last := (*h)[0]
		if intervals[i][0] >= last {
			// meeting can be held in the same room
			heap.Pop(h)
		}
		// add new meeting
		heap.Push(h, intervals[i][1])
	}

	return h.Len()
}

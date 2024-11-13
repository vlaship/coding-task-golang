package problem0253

import (
	"container/heap"
	"sort"
)

/*
253. Meeting Rooms II

https://leetcode.com/problems/meeting-rooms-ii/

Given an array of meeting time intervals intervals where intervals[i] = [starti, endi], return the minimum number of conference rooms required.

Example 1:

	Input: intervals = [[0,30],[5,10],[15,20]]
	Output: 2

Example 2:

	Input: intervals = [[7,10],[2,4]]
	Output: 1

Constraints:

	1 <= intervals.length <= 10^4
	0 <= starti < endi <= 10^6
*/

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

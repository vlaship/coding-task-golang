package problem2054

import "sort"

func maxTwoEvents(events [][]int) int {
	// Sort events by their start time
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	numOfEvents := len(events)

	// 'maxValueAfter' array will store the maximum value from current event to the last event
	maxValueAfter := make([]int, numOfEvents+1)
	for i := numOfEvents - 1; i >= 0; i-- {
		maxValueAfter[i] = max(maxValueAfter[i+1], events[i][2])
	}

	maxTotalValue := 0

	for _, event := range events {
		value := event[2] // Value of the current event

		// Binary search to find the first event that starts after the current event ends
		left, right := 0, numOfEvents
		for left < right {
			mid := (left + right) / 2
			if events[mid][0] > event[1] {
				// If the event at 'mid' starts after current event ends, search in left half
				right = mid
			} else {
				// Otherwise search in the right half
				left = mid + 1
			}
		}

		// If there is an event that starts after the current one, add its value
		if left < numOfEvents {
			value += maxValueAfter[left]
		}

		// Update the maximum total value if needed
		maxTotalValue = max(maxTotalValue, value)
	}

	return maxTotalValue
}

package problem1760

import "slices"

func minimumSize(nums []int, maxOperations int) int {
	left, right := 1, slices.Max(nums)

	for left < right {
		mid := (left + right) / 2

		operations := 0
		for _, num := range nums {
			if num > mid {
				operations += (num - 1) / mid
			}
		}

		if operations > maxOperations {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

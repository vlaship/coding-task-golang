package problem_0121

import "math"

func maxProfit(prices []int) int {
	buy := math.MaxInt
	sell := 0
	for _, price := range prices {
		if price < buy {
			buy = price
		}
		d := price - buy
		if d > sell {
			sell = d
		}
	}
	return sell
}

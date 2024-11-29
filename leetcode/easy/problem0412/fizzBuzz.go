package problem0412

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i < n+1; i++ {
		r := ""
		if i%3 == 0 {
			r += "Fizz"
		}
		if i%5 == 0 {
			r += "Buzz"
		}
		if r == "" {
			res[i-1] = strconv.Itoa(i)
		} else {
			res[i-1] = r
		}
	}
	return res
}

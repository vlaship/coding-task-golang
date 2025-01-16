package problem_0509

func fib(n int) int {
	if n == 0 {
		return 0
	}

	a, b := 0, 1
	for i := 2; i < n+1; i++ {
		t := a + b
		a = b
		b = t
	}

	return b
}

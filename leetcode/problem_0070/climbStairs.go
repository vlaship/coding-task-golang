package problem_0070

func climbStairs(n int) int {
	p1 := 1
	p2 := 1
	for i := 0; i < n-1; i++ {
		s := p1 + p2
		p1 = p2
		p2 = s
	}
	return p2
}

package countingbits

func countBits(n int) []int {
	res := []int{}
	for i := 0; i <= n; i++ {
		num := 0
		for t := i; t > 0; {
			num += t % 2
			t /= 2
		}
		res = append(res, num)
	}

	return res
}

package plusone

func plusOne(digits []int) []int {
	ans := make([]int, len(digits)+1)
	d := 1

	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + d
		if tmp <= 9 {
			d = 0
			ans[i+1] = tmp
			continue
		}

		d = 1
		ans[i+1] = 0
	}

	if d == 1 {
		ans[0] = 1
		return ans
	}

	return ans[1:]
}

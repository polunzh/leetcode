package perfectnumber

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1

	first := 2
	last := -1

	for {
		if first >= num/first {
			break
		}

		if num%first != 0 {
			first++
			continue
		}

		last = num / first

		sum += (first + last)

		first++
	}

	return sum == num
}

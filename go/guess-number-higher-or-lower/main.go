package guessNumber

const NUM int = 6

func guess(n int) int {
	if n < NUM {
		return 1
	}

	if n > NUM {
		return -1
	}

	return 0
}

func guessNumber(n int) int {
	low := 1
	high := n

	for {
		res := guess(n)

		if res == 0 {
			return n
		}

		if res == -1 {
			high = n
			n = low + (high - low) / 2
		} else {
			low = n
			n = n + (high-low)/2
		}
	}
}

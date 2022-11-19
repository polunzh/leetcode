package findthehighestaltitude

func largestAltitude(gain []int) int {
	max := 0
	pre := 0
	for i := 0; i < len(gain); i++ {
		cur := pre + gain[i]
		if cur > max {
			max = cur
		}

		pre = cur
	}

	return max
}

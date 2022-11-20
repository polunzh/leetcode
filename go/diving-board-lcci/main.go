package divingboardlcci

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}

	if shorter == longer {
		return []int{k}
	}

	var ans []int
	for i := k; i >= 0; i-- {
		ans = append(ans, i*shorter+(k-i)*longer)
	}

	return ans
}

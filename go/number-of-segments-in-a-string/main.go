package numberofsegmentsinastring

func countSegments(s string) int {
	res := 0

	preIndex := -1
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if preIndex == -1 {
				res += 1
			}
			preIndex = i
		} else {
			preIndex = -1
		}
	}

	return res
}

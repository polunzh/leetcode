package detectcapital

func detectCapitalUse(word string) bool {
	if len(word) <= 1 {
		return true
	}

	if len(word) == 2 {
		if word[0] >= 97 && word[1] <= 90 {
			return false
		}

		return true
	}

	first := word[0]
	second := word[1]

	var left byte = 97
	var right byte = 122
	if first <= 90 && second <= 90 {
		left = 65
		right = 90
	}

	for i := 1; i < len(word); i++ {
		if !(word[i] >= left && word[i] <= right) {
			return false
		}
	}

	return true
}

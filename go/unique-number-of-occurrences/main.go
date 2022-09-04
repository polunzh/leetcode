package uniquenumberofoccurrences

func uniqueOccurrences(arr []int) bool {
	freqMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if val, ok := freqMap[arr[i]]; ok {
			freqMap[arr[i]] = val + 1
			continue
		}

		freqMap[arr[i]] = 1
	}

	countMap := make(map[int]bool)

	for _, v := range freqMap {
		if _, ok := countMap[v];ok {
			return false
		}

		countMap[v] = true
	}

	return true
}

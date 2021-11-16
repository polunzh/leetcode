package ransomNote

func canConstruct(ransomNote string, magazine string) bool {
	magazinMap := make(map[rune]int, 26)
	for _, s := range magazine {
		magazinMap[s]++
	}

	for _, s := range ransomNote {
		value := magazinMap[s]
		if value == 0 {
			return false
		}

		magazinMap[s]--
	}

	return true
}

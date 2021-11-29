package KeyboardRow

var keyboardMap map[byte]int = make(map[byte]int)

func init() {
	var rows []string = []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	for index, row := range rows {
		for _, c := range row {
			keyboardMap[byte(c)] = index
			keyboardMap[byte(c)-32] = index
		}
	}
}

func findWords(words []string) []string {
	result := []string{}

	for _, word := range words {
		legal := true
		index := keyboardMap[word[0]]
		for _, w := range word {
			if keyboardMap[byte(w)] != index {
				legal = false
			}
		}

		if legal {
			result = append(result, word)
		}
	}

	return result
}

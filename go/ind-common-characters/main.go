package indCommonCharacters

import (
	"math"
)

func commonChars(words []string) []string {
	counter := make(map[rune]int, 26)    
	result := []string{}

	var i rune
	for i = 0; i < 26; i++ {
		counter[i] = 1000000
	}

	for _, w := range words {
		cnt := make(map[rune]int, 26)

		for _, c := range w {
			cnt[c - 'a']++
		}

		for i = 0; i < 26; i++ {
			counter[i] = int(math.Min(float64(counter[i]), float64(cnt[i])))
		}
	}

	for i := range counter {
		for j := 0; j < counter[i]; j++ {
			result = append(result, string(i + 'a'))
		}
	}

	return result
}
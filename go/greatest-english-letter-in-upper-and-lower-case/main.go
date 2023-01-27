// https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/

package greatestenglishletterinupperandlowercase

func greatestLetter(s string) string {
	var answer byte
	charMap := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		c := s[i]
		upper := s[i]
		if c >= 65 && c <= 90 {
			c += 32
		} else {
			c -= 32
			upper = c
		}

		if ok, _ := charMap[c]; ok && upper > answer {
			answer = upper
		}

		charMap[s[i]] = true
	}

	if answer == 0 {
		return ""
	}

	return string(answer)
}

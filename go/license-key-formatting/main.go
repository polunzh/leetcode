package licenseKeyFormatting

import "strings"

func licenseKeyFormatting(s string, k int) string {
	if strings.Trim(s, "") == "" {
		return s
	}

	result := ""
	counter := 0

	diff := byte('a' - 'A')

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}

		c := s[i]
		if c >= 'a' && c <= 'z' {
			c = s[i] - diff
		}

		result = string(c) + result
		counter++
		if counter == k && i != 0 {
			result = "-" + result
			counter = 0
		}
	}

	if len(result) == 0 {
		return result
	}

	if result[0] == '-' {
		return result[1:]
	}

	return result
}

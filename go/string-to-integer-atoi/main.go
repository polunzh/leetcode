package stringtointegeratoi

import (
	"math"
)

func myAtoi(s string) int {
	res := 0
	symbol := ""

	i := 0

	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}

	hasZeroPrefix := false

	for ; i < len(s); i++ {
		if hasZeroPrefix && (s[i] == '+' || s[i] == '-') {
			return 0
		}

		if s[i] != '0' {
			break
		}

		hasZeroPrefix = true
	}

	if i == len(s) {
		return 0
	}

	if s[i] != '+' && s[i] != '-' && (s[i] <= '0' || s[i] > '9') {
		return 0
	}

	for ; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' {
			if res != 0 {
				break
			}

			if symbol != "" {
				return 0
			}

			symbol = string(s[i])
			continue
		}

		v := s[i] - '0'
		if v < 0 || v > 9 {
			break
		}

		res = res*10 + int(v)
		if (symbol == "" || symbol == "+") && res > math.MaxInt32 {
			return math.MaxInt32
		}

		if symbol == "-" && res*(-1) < math.MinInt32 {
			return math.MinInt32
		}
	}

	if symbol == "-" {
		return res * (-1)
	}

	return res
}

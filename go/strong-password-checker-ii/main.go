package strongpasswordcheckerii

import (
	"strings"
)

func strongPasswordCheckerII(password string) bool {
	password = strings.Trim(password, " ")

	if password == "" {
		return false
	}

	if len(password) < 8 {
		return false
	}

	hasLowerCase := false
	hasUpperCase := false
	hasDigit := false
	specialChar := false

	var lastChar byte

	var SpecialCharMap = map[byte]bool{
		'!': true,
		'@': true,
		'#': true,
		'$': true,
		'%': true,
		'^': true,
		'&': true,
		'*': true,
		'(': true,
		')': true,
		'-': true,
		'+': true,
	}

	for i := 0; i < len(password); i++ {
		if lastChar == password[i] {
			return false
		}

		v := password[i]
		lastChar = v

		if v >= 'a' && v <= 'z' {
			hasLowerCase = true
			continue
		}

		if v >= 'A' && v <= 'Z' {
			hasUpperCase = true
			continue
		}

		if v >= '0' && v <= '9' {
			hasDigit = true
			continue
		}

		if ok, _ := SpecialCharMap[v]; ok {
			specialChar = true
		}
	}

	return hasLowerCase && hasUpperCase && hasDigit && specialChar
}

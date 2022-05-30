package validparentheses

func isValid(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 {
			return false
		}

		stackLen := len(stack)
		last := stack[stackLen-1]

		if (last == '{' && s[i] == '}') || (last == '(' && s[i] == ')') || (last == '[' && s[i] == ']') {
			stack = stack[0 : stackLen-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

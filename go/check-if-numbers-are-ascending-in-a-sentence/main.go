package checkifnumbersareascendinginasentence

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}

func areNumbersAscending(s string) bool {
	preNum := -1
	curNum := 0

	for i := 0; i < len(s); i++ {
		if !isNum(s[i]) {
			if curNum <= preNum && curNum > 0 {
				return false
			}

			if curNum > 0 {
				preNum = curNum
				curNum = 0
			}
			continue
		}

		curNum = curNum*10 + int(s[i]-'0')
	}

	if curNum != 0 {
		return preNum < curNum
	}

	return true
}

package numberComplement

import (
	"fmt"
	"math"
)

func toBinary(num int) string {
	if num <= 0 {
		return "0"
	}

	res := ""
	for num > 0 {
		res += fmt.Sprintf("%d", num%2)
		num /= 2
	}

	return res
}

func findComplement(num int) int {
	var result int
	binary := toBinary(num)
	for index, n := range binary {
		if n == '0' {
			result += int(math.Pow(2, float64(index)))
		}
	}

	return result
}

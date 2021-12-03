package bitwiseComplement

import (
	"fmt"
	"math"
)

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	bin := ""

	for n > 0 {
		var v int
		if n%2 == 0 {
			v = 1
		} else {
			v = 0
		}

		bin = fmt.Sprintf("%d%s", v, bin)
		n /= 2
	}

	res := 0
	for i := 0; i < len(bin); i++ {
		res += (int(bin[i] - '0')) * int(math.Pow(float64(2), float64(len(bin)-i-1)))
	}

	return res
}

package main

import (
	"fmt"
)

func isPowerOfFour(n int) bool {
	for res := 1; res <= n; res *= 4 {
		if res == n {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(isPowerOfFour(1))
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(5))
}

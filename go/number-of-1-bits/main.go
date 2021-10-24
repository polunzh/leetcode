package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	var counter int = 0
	for num != 0 {
		if num & 1 == 1 {
			counter++
		}
		num >>= 1
	}

	return counter
}

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
	fmt.Println(hammingWeight(00000000000000000000000010000000))
}

package main

import "fmt"

func isUgly(n int) bool {
	for _, num := range [3]int{2, 3, 5} {
		for n != 0 && n%num == 0 {
			n = n / num
		}
	}

	return n == 1
}

func main() {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(8))
	fmt.Println(isUgly(14))
	fmt.Println(isUgly(1))
}

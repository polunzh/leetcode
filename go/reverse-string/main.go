package main

import "fmt"

func reverseString(s []byte) {
	var length = len(s)
	var mid int = len(s) / 2
	for i := 0; i < mid; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

func main() {
	s := []byte{'z', 'h', 'a', 'n', 'g'}
	fmt.Println(s)
	reverseString(s)
	fmt.Println(s)
}

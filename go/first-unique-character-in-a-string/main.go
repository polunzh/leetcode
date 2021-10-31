package main

import "fmt"

func firstUniqChar(s string) int {
	var feq [26]int
	for i := 0; i < len(s); i++ {
		feq[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if feq[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}

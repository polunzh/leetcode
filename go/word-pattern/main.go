package main

import "fmt"

func wordPattern(pattern string, s string) bool {
	patternMap := make(map[byte]string)
	wordMap := make(map[string]byte)
	var index int = 0
	for i := 0; i < len(pattern); i++ {
		word := ""
		for ; index < len(s) && s[index] != ' '; index++ {
			word += string(s[index])
		}

		// 单词个数和 pattern 长度不一样
		if (i+1 == len(pattern) && index != len(s)) || (i+1 != len(pattern) && index == len(s)) {
			return false
		}

		index++
		if patternMap[pattern[i]] != "" || wordMap[word] != byte(0) {
			if patternMap[pattern[i]] != word || wordMap[word] != pattern[i] {
				return false
			}
		} else {
			patternMap[pattern[i]] = word
			wordMap[word] = pattern[i]
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abbc", "dog cat cat fish"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
	fmt.Println(wordPattern("abc", "dog cat dog"))
	fmt.Println(wordPattern("aaa", "aa aa aa aa"))
}

/*
* 参考这里的答案实现的: https://leetcode-cn.com/problems/number-of-matching-subsequences/solution/pi-pei-zi-xu-lie-de-dan-ci-shu-by-leetcode/
*/
package numberOfMatchingSubsequences

type node struct {
	word  string
	index int
}

func NumMatchingSubseq(s string, words []string) int {
	buckets := make(map[byte][]node, 26)
	result := 0

	for _, word := range words {
		buckets[word[0]-'a'] = append(buckets[word[0]-'a'], node{word, 0})
	}

	for _, char := range s {
		oldBucket := buckets[byte(char)-'a']
		buckets[byte(char)-'a'] = []node{}

		for _, n := range oldBucket {
			n.index++
			if n.index == len(n.word) {
				result++
			} else {
				buckets[n.word[n.index]-'a'] = append(buckets[n.word[n.index]-'a'], node{n.word, n.index})
			}
		}
	}

	return result
}

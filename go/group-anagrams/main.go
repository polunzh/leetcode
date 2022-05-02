package groupanagrams

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		t := SortString(strs[i])
		if _, ok := strMap[t]; ok {
			strMap[t] = append(strMap[t], strs[i])
		} else {
			strMap[t] = []string{strs[i]}
		}
	}

	result := [][]string{}

	for _, value := range strMap {
		result = append(result, value)
	}

	return result
}

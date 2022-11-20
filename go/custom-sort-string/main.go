package customsortstring

import "sort"

func customSortString(order string, s string) string {
	h := [26]int{}
	vis := [26]bool{}
	idx := 26

	for i := 0; i < len(order); i++ {
		h[order[i]-'a'] = idx
		idx--
		vis[order[i]-'a'] = true
	}

	for c := 'a'; c <= 'z'; c++ {
		if !vis[c-'a'] {
			h[c-'a'] = idx
			idx--
		}
	}

	ans := []byte(s)

	sort.Slice(ans, func(i, j int) bool {
		return h[ans[i]-'a'] > h[ans[j]-'a']
	})

	return string(ans)
}

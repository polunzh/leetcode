package IsSubsequence

func IsSubsequence(s string, t string) bool {
	var indexS, indexT int
	sLen := len(s)
	tLen := len(t)
	for ; indexS < sLen && indexT < tLen; indexT++ {
		if s[indexS] == t[indexT] {
			indexS++
		}
	}

	return indexS == sLen
}

package reshapethematrix

func matrixReshape(mat [][]int, r int, c int) [][]int {
	count := len(mat) * len(mat[0])
	if count != r*c {
		return mat
	}

	var arr []int
	for _, a := range mat {
		arr = append(arr, a...)
	}

	var res [][]int = make([][]int, r)
	 c1 := len(mat[0])
	for i := 0; i < count; i++ {
		if len(res[i/c]) == 0 {
			res[i/c] = make([]int, c)
		}
		res[i/c][i%c] = mat[i/c1][i%c1]
	}

	return res
}

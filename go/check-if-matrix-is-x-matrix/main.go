package checkifmatrixisxmatrix

func checkXMatrix(grid [][]int) bool {
	length := len(grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// 对角线
			if i == j || ((i + j) == length-1) {
				if grid[i][j] == 0 {
					return false
				}

				continue
			}

			if grid[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

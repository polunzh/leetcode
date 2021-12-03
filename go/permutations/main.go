package permutation

func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	var path []int

	dfs(nums, 0, path, used, &res)
	return res
}

func dfs(nums []int, depth int, path []int, used []bool, res *([][]int)) {
	if len(nums) == depth {
		pathClone := make([]int, len(path))
		copy(pathClone, path)
		*res = append(*res, pathClone)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] == true {
			continue
		}

		path = append(path, nums[i])
		used[i] = true
		dfs(nums, depth+1, path[:], used, res)
		used[i] = false
		path = path[0 : len(path)-1]
	}
}

package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	paths := [][]int{}
	find(candidates, target, []int{}, &paths)
	return paths
}

func find(candidates []int, target int, path []int, paths *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		*paths = append(*paths, path)
		return
	}

	for i, v := range candidates {
		var newPath []int = make([]int, len(path))
		copy(newPath, path)
		newPath = append(newPath, v)

		find(candidates[i:], target-v, newPath, paths)
	}
}

func main() {
	var target int = 8
	var candidates []int = make([]int, 3)
	candidates[0] = 2
	candidates[1] = 3
	candidates[2] = 5

	fmt.Println(combinationSum(candidates, target))
}

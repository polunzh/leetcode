package relativeSortArray

func max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	result := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > result {
			result = arr[i]
		}
	}

	return result
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	result := []int{}
	maxValue := max(arr1)
	frequency := make([]int, maxValue+1)

	for i := 0; i < len(arr1); i++ {
		frequency[arr1[i]]++
	}

	for i := 0; i < len(arr2); i++ {
		for j := 0; j < frequency[arr2[i]]; j++ {
			result = append(result, arr2[i])
		}

		frequency[arr2[i]] = 0
	}

	for i := 0; i < len(frequency); i++ {
		for j := 0; j < frequency[i]; j++ {
			result = append(result, i)
		}
	}

	return result
}

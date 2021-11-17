package findAllNumbersDisappearedInAnArray

func findDisappearedNumbers(nums []int) []int {
	table := make([]int, len(nums))
	for _, v := range nums {
		table[v-1] = 1
	}

	result := []int{}
	for i, v := range table {
		if v == 0 {
			result = append(result, i+1)
		}
	}

	return result
}

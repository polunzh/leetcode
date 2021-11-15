package twoSum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for index, n := range nums {
		if v, exist := m[target-n]; exist {
			return []int{v, index}
		}

		m[n] = index
	}

	return []int{}
}

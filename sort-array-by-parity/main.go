package sortarraybyparity

func sortArrayByParity(nums []int) []int {
	if nums == nil {
		return nums
	}

	splitor := -1

	for index := 0; index < len(nums); index++ {
		isEven := nums[index]%2 == 0
		if isEven && splitor != -1 {
			tmp := nums[splitor]
			nums[splitor] = nums[index]
			nums[index] = tmp
			splitor++
		}

		if !isEven && splitor == -1 {
			splitor = index
		}
	}

	return nums
}

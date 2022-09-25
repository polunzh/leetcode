package sortevenandoddindicesindependently

import "sort"

func sortEvenOdd(nums []int) []int {
	var evenArr, oddArr []int
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			evenArr = append(evenArr, nums[i])
		} else {
			oddArr = append(oddArr, nums[i])
		}
	}

	sort.Ints(evenArr)
	sort.Sort(sort.Reverse(sort.IntSlice(oddArr)))

	ans := make([]int, len(nums))
	for i := 0; i < len(evenArr); i++ {
		ans[2*i] = evenArr[i]
	}

	for i := 0; i < len(oddArr); i++ {
		ans[2*i+1] = oddArr[i]
	}

	return ans
}

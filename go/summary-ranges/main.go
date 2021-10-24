package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}

	var result []string = []string{}
	var start int = nums[0]

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > 1 {
			if start == nums[i] {
				result = append(result, fmt.Sprintf("%d", nums[i]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i]))
			}
			start = nums[i+1]
		}

		if i == len(nums)-2 {
			if start == nums[i+1] {
				result = append(result, fmt.Sprintf("%d", nums[i+1]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i+1]))
			}
		}
	}

	return result
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{7}))
	fmt.Println(summaryRanges([]int{0, 1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

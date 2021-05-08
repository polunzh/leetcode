package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	var v1 int = math.MinInt64
	var v2 int = math.MinInt64
	var v3 int = math.MinInt64

	for index := 0; index < len(nums); index++ {
		if nums[index] == v1 || nums[index] == v2 || nums[index] == v3 {
			continue
		}

		if v1 == math.MinInt64 || nums[index] > v1 {
			v3 = v2
			v2 = v1
			v1 = nums[index]
		} else if v2 == math.MinInt64 || nums[index] > v2 {
			v3 = v2
			v2 = nums[index]
		} else if v3 == math.MinInt64 || nums[index] > v3 {
			v3 = nums[index]
		}
	}

	if v3 == math.MinInt64 {
		return v1
	}

	return v3
}

func main() {
	var nums []int

	nums = []int{2, 3, 1}
	fmt.Println(nums, thirdMax(nums))

	nums = []int{2, 1}
	fmt.Println(nums, thirdMax(nums))

	nums = []int{2, 1, 1}
	fmt.Println(nums, thirdMax(nums))

	nums = []int{1, 2, 2}
	fmt.Println(nums, thirdMax(nums))

	nums = []int{1, 2, 2, 5, 3, 5}
	fmt.Println(nums, thirdMax(nums))

	nums = []int{1}
	fmt.Println(nums, thirdMax(nums))

	nums = []int{2, 2, 3, 1}
	fmt.Println(nums, thirdMax(nums))

	nums = []int{1, 1, 1, 1}
	fmt.Println(nums, thirdMax(nums))

	nums = []int{-1, 1, 31, 1}
	fmt.Println(nums, thirdMax(nums))

	nums = []int{5, 2, 4, 1, 3, 6, 0}
	fmt.Println(nums, thirdMax(nums))
}

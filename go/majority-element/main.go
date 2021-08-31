package main

import "fmt"

func majorityElement(nums []int) int {
	var length = len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	counter := make(map[int]int, length)

	for _, v := range nums {
		value, exists := counter[v]
		if exists == true {
			if value >= length/2 {
				return v
			}

			counter[v] = value + 1
		} else {

			counter[v] = 1
		}
	}

	return 0
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

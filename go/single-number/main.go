package main

import "fmt"

func main() {
	var nums []int = []int{2, 2, 1}
	var nums2 []int = []int{4,1,2,1,2}
	fmt.Println(singleNumber(nums))	
	fmt.Println(singleNumber(nums2))	
}

func singleNumber(nums []int) int {
	var res int = 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res;
}

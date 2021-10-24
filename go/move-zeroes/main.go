package main

import "fmt"

func moveZeroes(nums []int) {
	var zeroCounter int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCounter++
		} else if zeroCounter != 0 {
			nums[i-zeroCounter], nums[i] = nums[i], nums[i-zeroCounter]
		}
	}
}

func main() {
	var nums = []int{0, 1, 3, 12, 0, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

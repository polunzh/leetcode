package main

import "fmt"

type subArray struct {
	start     int
	end       int
	frequency int
}

func findShortestSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := 0
	maxFrequency := 0
	m := make(map[int]subArray)

	for index, n := range nums {
		temp, ok := m[n]
		if !ok {
			m[n] = subArray{start: index, end: index, frequency: 1}
		} else {
			temp.end = index
			temp.frequency++
			m[n] = temp
		}

		temp, _ = m[n]
		currentDegree := temp.end - temp.start + 1
		if temp.frequency > maxFrequency {
			maxFrequency = temp.frequency
			result = currentDegree
		} else if temp.frequency == maxFrequency && currentDegree < result {
			maxFrequency = temp.frequency
			result = currentDegree
		}
	}

	return result
}

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{8, 9, 8, 9, 8}))
	fmt.Println(findShortestSubArray([]int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2}))
}

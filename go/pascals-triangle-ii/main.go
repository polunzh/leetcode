package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	}

	if rowIndex == 0 {
		return []int{1}
	}

	result := []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		temp := []int{1}
		for t := 1; t < len(result); t++ {
			temp = append(temp, result[t-1]+result[t])
		}

		result = append(temp, 1)
	}

	return result
}

func main() {
	fmt.Println(getRow(1))
	fmt.Println(getRow(2))
	fmt.Println(getRow(3))
	fmt.Println(getRow(4))
}

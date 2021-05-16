package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {
	var l, w int
	var max int = 0
	var result []int

	var limit = int(math.Sqrt(float64(area)))

	for w = 1; w <= limit; w++ {
		if (area % w) == 0 {

			l = area / w
			if w > l {
				break
			}

			if max == 0 {
				max = l + w
				result = []int{l, w}

				continue
			}

			if (l + w) < max {
				max = l + w
				result = []int{l, w}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(4, constructRectangle(4))
	fmt.Println(37, constructRectangle(37))
	fmt.Println(122122, constructRectangle(122122))
}

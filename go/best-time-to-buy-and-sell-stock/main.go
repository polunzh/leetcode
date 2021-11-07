package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	baseNumber := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		tempProfit := prices[i] - baseNumber
		if prices[i] > baseNumber && tempProfit > profit {
			profit = tempProfit
		} else if prices[i] < baseNumber {
			baseNumber = prices[i]
		}
	}

	return profit
}

func main() {
	var stock []int
	stock = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(stock))

	stock = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(stock))
}

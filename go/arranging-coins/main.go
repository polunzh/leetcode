package arrangingcoins

// 该题目可以使用二分查找的方法
func arrangeCoins(n int) int {
	if n < 1 {
		return 0
	}

	minLayer, maxLayer := 1, n

	for minLayer < maxLayer {
		mid := minLayer + (maxLayer - minLayer + 1) / 2
		if mid*(mid+1) <= 2*n {
			minLayer = mid
		} else {
			maxLayer = mid - 1
		}
	}

	return minLayer
}

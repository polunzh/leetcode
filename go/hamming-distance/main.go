package hammingDistance

func hammingDistance(x int, y int) int {
	distance := 0

	for x > 0 || y > 0 {
		x1 := x % 2
		y1 := y % 2
		if x1 != y1 {
			distance++
		}

		x /= 2
		y /= 2
	}

	return distance
}

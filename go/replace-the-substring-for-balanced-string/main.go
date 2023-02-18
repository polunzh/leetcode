package replacethesubstringforbalancedstring

func findTheArrayConcVal(nums []int) int64 {
	var sum int = 0

	i := 0
	j := len(nums) - 1

	for i <= j {
		if i == j {
			sum += nums[i]
			break
		}

		t := 1
		v := nums[j]
		for v > 0 {
			v /= 10
			t *= 10
		}

		sum += (nums[i]*t + nums[j])
		i++
		j--
	}

	return int64(sum)
}

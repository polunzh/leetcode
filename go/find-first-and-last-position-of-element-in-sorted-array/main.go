package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := binarySearch(nums, 0, target)

	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	right := binarySearch(nums, left+1, target+1)

	return []int{left, right - 1}
}

func binarySearch(nums []int, start int, target int) int {
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if target > nums[mid] {
			start = mid + 1
			continue
		}

		end = mid - 1
	}

	return start
}

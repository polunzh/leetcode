package intersectionOfTwoArrays

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]] = true
	}

	for j := 0; j < len(nums2); j++ {
		m2[nums2[j]] = true
	}

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}

	var result []int
	for item := range m1 {
		if _, ok := m2[item]; ok {
			result = append(result, item)
		}
	}

	return result
}

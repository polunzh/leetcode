package intersectionOfTwoArraysII

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]] ++
	}

	for j := 0; j < len(nums2); j++ {
		m2[nums2[j]] ++
	}

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}

	var result []int
	for item := range m1 {
		if _, ok := m2[item]; ok {
			l1 := m1[item]
			l2 := m2[item]
			l := l1
			if l1 >= l2 {
				l = l2
			}

			for x := 0; x < l; x++ {
				result = append(result, item)
			}
		}
	}

	return result
}

package relativeranks

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	var result []string

	var clonedScore []int = make([]int, len(score))
	for i := 0; i < len(score); i++ {
		clonedScore[i] = score[i]
	}

	sort.Slice(clonedScore, func(i, j int) bool {
		return clonedScore[i] >= clonedScore[j]
	})

	scoreMap := make(map[int]string, len(clonedScore))
	medals := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}

	for i := 0; i < len(clonedScore); i++ {
		if i < 3 {
			scoreMap[clonedScore[i]] = medals[i]
			continue
		}

		scoreMap[clonedScore[i]] = strconv.Itoa(i + 1)
	}

	for i := 0; i < len(score); i++ {
		result = append(result, scoreMap[score[i]])
	}

	return result
}

package numberofseniorcitizens

import (
	"strconv"
)

func countSeniors(details []string) int {
	count := 0
	for _, v := range details {

		if v, err := strconv.Atoi(v[11:13]); err == nil && v > 60 {
			count++
		}
	}


	return count
}

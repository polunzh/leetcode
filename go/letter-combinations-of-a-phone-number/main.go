package lettercombinationsofaphonenumber

import "fmt"

var KEY_MAP = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var result []string = []string{}
	for _, v := range digits {
		arr, ok := KEY_MAP[byte(v)]
		if !ok {
			panic("Invalid digit")
		}

		t := []string{}
		if len(result) == 0 {
			result = arr
			continue
		}

		for _, k := range arr {
			for _, j := range result {
				t = append(t, fmt.Sprintf("%s%s", string(j), string(k)))
			}
		}
		result = t
	}

	return result
}

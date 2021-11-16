package addStrings

import "fmt"

func addStrings(num1 string, num2 string) string {
	result := ""

	index1 := len(num1) - 1
	index2 := len(num2) - 1
	var r byte = 0
	for index1 >= 0 && index2 >= 0 {
		t := (num1[index1] - '0') + (num2[index2] - '0') + r
		if t > 9 {
			r = t / 10
			t = t % 10
		} else {
			r = 0
		}

		result = fmt.Sprintf("%d", t) + result
		index1--
		index2--
	}

	num := ""
	index := -1
	if index1 > -1 {
		num = num1
		index = index1
	} else if index2 > -1 {
		num = num2
		index = index2
	} else if r > 0 {
		result = fmt.Sprintf("%d", r) + result
		r = 0
	}

	for index >= 0 {
		t := (num[index] - '0') + r
		if t > 9 {
			r = t / 10
			t = t % 10
		} else {
			r = 0
		}

		result = fmt.Sprintf("%d", t) + result
		index--
	}

	if r > 0 {
		result = fmt.Sprintf("%d", r) + result
		r = 0
	}

	return result
}

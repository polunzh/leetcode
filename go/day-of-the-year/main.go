package dayOfYearLib

import (
	"strconv"
)

func isLeap(year int) bool {
	if year%4 == 0 || year%400 == 0 {
		return true
	}

	return false
}

func dayOfYear(date string) int {
	daysOfMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, _ := strconv.Atoi(date[0:4])
	if isLeap(year) {
		daysOfMonth[1] = 29
	}

	month, _ := strconv.Atoi(date[5:7])
	result := 0
	for i := 0; i < month-1; i++ {
		result += daysOfMonth[i]
	}

	day, _ := strconv.Atoi(date[8:])
	result += day

	return result
}

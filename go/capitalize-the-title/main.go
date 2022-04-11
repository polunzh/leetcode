package capitalizethetitle

import "strings"

func capitalizeTitle(title string) string {
	var arr []string = make([]string, len(title))

	for i, j := 0, 0; i <= len(title); i++ {
		if i == len(title) || title[i] == ' ' {
			if i-j > 2 {
				arr[j] = strings.ToUpper(string(title[j]))
			}

			if i == len(title) {
				break
			}

			arr[i] = string(title[i])
			j = i + 1
		} else {
			arr[i] = strings.ToLower(string(title[i]))
		}
	}

	return strings.Join(arr, "")
}

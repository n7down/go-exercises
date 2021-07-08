package numberchars

import "fmt"

func NumberChars(input string) string {
	var r string
	var last string = ""
	var count int = 1
	for _, c := range input {
		if string(c) != last {
			r += fmt.Sprintf("%v%v", count, string(c))
			count = 1
			last = string(c)
		} else {
			count = count + 1
		}
	}
	return r
}

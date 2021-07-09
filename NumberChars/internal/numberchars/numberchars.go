package numberchars

import "fmt"

func NumberChars(input string) string {
	var r string
	var last string = ""
	var count int = 0
	for _, c := range input {
		if last == "" {
			count = count + 1
			last = string(c)
		} else if string(c) != last {
			r += fmt.Sprintf("%v%v", count, string(last))
			count = 1
			last = string(c)
		} else {
			count = count + 1
		}
	}
	r += fmt.Sprintf("%v%v", count, string(last))
	return r
}

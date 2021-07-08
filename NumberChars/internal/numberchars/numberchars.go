package numberchars

import "fmt"

func NumberChars(input string) string {
	var r string
    var first string = string(input[0])
    var rest = string(input[1:])
	var last string = first
	var count int = 1
	for _, c := range rest {
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

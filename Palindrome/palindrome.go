package main

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	runes := []rune(s)
	first := string(runes[0 : len(s)/2])
	r := Reverse(s)
	runes = []rune(r)
	second := string(runes[0 : len(r)/2])
	if first == second {
		return true
	}
	return false
}

func main() {

}

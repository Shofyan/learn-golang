package main

func main() {

	println(IsPalindrome("asti"))
	println(IsPalindrome("kodok"))

	println(RecursiveIspalindrome("asti", len("asti")/2))
	println(RecursiveIspalindrome("kodok", len("kodok")/2))

}

func IsPalindrome(s string) bool {
	strRune := []rune(s)
	for i := len(s) / 2; i >= 0; i-- {
		if string(strRune[len(s)-1-i]) != string(strRune[i]) {
			return false
		}
	}
	return true
}

func RecursiveIspalindrome(s string, i int) bool {
	if i < 0 {
		return true
	} else {
		strRune := []rune(s)
		if string(strRune[len(s)-1-i]) != string(strRune[i]) {
			return false
		}
		return RecursiveIspalindrome(s, i-1)
	}

}

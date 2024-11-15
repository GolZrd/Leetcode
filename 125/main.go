package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		l := rune(s[left])
		r := rune(s[right])
		if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			left++
		} else if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			right--
		} else if unicode.ToLower(l) == unicode.ToLower(r) {
			left++
			right--
		} else {
			return false
		}

	}

	return true
}

// // Решение с использованием Strings.Map
// func isPalindrome(s string) bool {
// 	modified := func(r rune) rune {
// 		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
// 			return -1
// 		}
// 		return unicode.ToLower(r)
// 	}

// 	s = strings.Map(modified, s)

// 	l, r := 0, len(s)-1
// 	for l < r {
// 		if s[l] != s[r] {
// 			return false
// 		}
// 		l++
// 		r--
// 	}

// 	return true
// }

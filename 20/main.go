package main

import "fmt"

func main() {
	fmt.Println(isValid("){")) // ([])
}

func isValid(s string) bool {

	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}

	for _, i := range s {
		if k, ok := m[i]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != k {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, i)
		}
	}

	return len(stack) == 0
}

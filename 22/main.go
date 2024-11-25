package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	stack := []string{}
	res := []string{}

	var backtrack func(int, int)
	backtrack = func(open, close int) {
		if open == n && close == n {
			res = append(res, strings.Join(stack, ""))
			fmt.Println("Когда open == n && close == n стак равен ", stack)
			return
		}

		if open < n {
			stack = append(stack, "(")
			fmt.Println("До функции, когда open < n стак равен ", stack)
			backtrack(open+1, close)
			fmt.Println("После функции стак равен ", stack, " далее удаляем символ с конца")
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			fmt.Println("До функции, когда close < open стак равен ", stack)
			backtrack(open, close+1)
			fmt.Println("После функции стак равен ", stack, " далее удаляем символ с конца")
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return res
}

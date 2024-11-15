package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr("255.100.50.0"))
}

// Решение 1
func defangIPaddr(address string) string {
	var st strings.Builder

	for _, i := range address {
		if i == '.' {
			st.WriteString("[.]")
		} else {
			st.WriteRune(i)
		}
	}
	return st.String()
}

// Решение с помощью сторонней библиотеки
// func defangIPaddr(address string) string {
// 	return strings.ReplaceAll(address, ".", "[.]")
// }

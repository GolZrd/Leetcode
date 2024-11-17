package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mergeAlternately("ab", "pqrs"))
}

func mergeAlternately(word1 string, word2 string) string {
	st := strings.Builder{}

	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		st.WriteByte(word1[i])
		st.WriteByte(word2[j])
		i++
		j++
	}

	if i < len(word1) {
		st.WriteString(string(word1[i:]))
	}
	if j < len(word2) {
		st.WriteString(string(word2[j:]))
	}

	return st.String()

}

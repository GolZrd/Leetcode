package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}

// Первое решение
// func uniqueMorseRepresentations(words []string) int {
// 	morseMap := map[byte]string{
// 		'a': ".-",
// 		'b': "-...",
// 		'c': "-.-.",
// 		'd': "-..",
// 		'e': ".",
// 		'f': "..-.",
// 		'g': "--.",
// 		'h': "....",
// 		'i': "..",
// 		'j': ".---",
// 		'k': "-.-",
// 		'l': ".-..",
// 		'm': "--",
// 		'n': "-.",
// 		'o': "---",
// 		'p': ".--.",
// 		'q': "--.-",
// 		'r': ".-.",
// 		's': "...",
// 		't': "-",
// 		'u': "..-",
// 		'v': "...-",
// 		'w': ".--",
// 		'x': "-..-",
// 		'y': "-.--",
// 		'z': "--..",
// 	}

// 	count := 0
// 	morse := make(map[string]struct{})
// 	st := strings.Builder{}

// 	for _, i := range words {
// 		for _, j := range i {
// 			st.WriteString(morseMap[byte(j)])
// 		}
// 		if _, ok := morse[st.String()]; !ok {
// 			morse[st.String()] = struct{}{}
// 			count++
// 		}
// 		st.Reset()
// 	}

// 	return count
// }

// Более оптимальное решение
func uniqueMorseRepresentations(words []string) int {
	morseCode := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	uniqueWords := make(map[string]struct{})
	st := strings.Builder{}

	for _, word := range words {
		for _, char := range word {
			index := char - 'a'
			st.WriteString(morseCode[index])
		}
		if _, ok := uniqueWords[st.String()]; !ok {
			uniqueWords[st.String()] = struct{}{}
		}
		st.Reset()
	}

	return len(uniqueWords)
}

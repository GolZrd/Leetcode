package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	in := []string{"neet", "code", "love", "you"}

	s := &Solution{}

	fmt.Println(s.Encode(in))
	fmt.Println(s.Decode(s.Encode(in)))

}

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res := strings.Builder{}
	for _, i := range strs {
		res.WriteString(strconv.Itoa(len(i)) + "#" + i)
	}
	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	res, i := []string{}, 0

	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		leng, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		res = append(res, encoded[i:i+leng])
		i += leng
	}
	return res
}

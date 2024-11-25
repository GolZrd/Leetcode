package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}

func lengthOfLongestSubstring(s string) int {
	res := 0
	left := 0
	mp := make(map[rune]bool)
	for i, v := range s {
		for mp[v] {
			delete(mp, rune(s[left]))
			left++
		}
		mp[v] = true
		if i-left+1 > res {
			res = i - left + 1
		}
	}

	return res
}

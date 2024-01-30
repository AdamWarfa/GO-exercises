package exercises

import (
	"fmt"
	"strings"
)

func TestLongestSubstring() {
	s := "abcabcbb"
	s2 := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstring(s2))
}

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	longest := ""

	current := ""

	for right < len(s) {
		letter := string(s[right])
		if strings.Contains(s, letter) {
			left++
			right = left
			current = ""
		} else {
			current += letter
			fmt.Println("her", current)
			if len(current) > len(longest) {
				longest = current
			}
			right++
		}
	}
	fmt.Println(longest)
	return len(longest)
}

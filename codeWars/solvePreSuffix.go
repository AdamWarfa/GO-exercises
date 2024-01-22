package kata

import (
	"fmt"
	"strings"
)

func Solve(s string) int {
	var prefix string
	var suffix string

	for i := 0; i < 1; i++ {
		for j := 1; j < len(s)/2+1; j++ {
			prefix = s[i:j]
			if strings.Contains(s[len(s)-len(prefix):], prefix) {
				suffix = prefix
			}
		}
	}
	fmt.Println("answer is ", len(suffix))
	return len(suffix)
}

package kata

import (
	"fmt"
)

func Solution(str, ending string) bool {
	if len(ending) <= len(str) {
		fmt.Println(str[len(str)-len(ending):] == ending)
		return str[len(str)-len(ending):] == ending
	} else {
		fmt.Println(false)
		return false
	}
}

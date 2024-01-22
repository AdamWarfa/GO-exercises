package kata

import "fmt"

func PrinterError(s string) string {
	count := 0
	for _, letter := range s {
		if string(letter) > "m" {
			count++
		}
	}
	answer := fmt.Sprintf("%v/%v", count, len(s))

	return answer
}

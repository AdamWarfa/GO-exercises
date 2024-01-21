package exercises

import (
	"slices"
	"strings"
)

/*
Write a function that reverses a given string.
*/

func ReverseString(str string) string {

	stringList := strings.Split(str, "")

	slices.Reverse(stringList)

	return strings.Join(stringList, "")
}

// func ReverseString(str string) string {
//     runes := []rune(str)
//     for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//         runes[i], runes[j] = runes[j], runes[i]
//     }
//     return string(runes)
// }

package exercises

import "slices"

/*
Write a function that takes a slice of integers
and returns a new slice containing only the unique elements.
*/

func UniqueElements(numbers []int) []int {
	result := []int{}
	for _, num := range numbers {
		if slices.Contains(result, num) {
			continue
		} else {
			result = append(result, num)
		}
	}

	return result
}

// func UniqueElements(numbers []int) []int {
// 	uniqueSet := make(map[int]bool)
// 	result := []int{}

// 	for _, num := range numbers {
// 		if !uniqueSe[num] {
// 			uniqueSet[num] = true
// 			result = append(result, num)
// 		}
// 	}

// 	return result
// }

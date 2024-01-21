package exercises

/*
Write a program that takes a slice of integers
and creates two slices, one containing even numbers
and the other containing odd numbers.
*/

func SeparateEvenOdd(numbers []int) ([]int, []int) {
	even := []int{}
	odd := []int{}

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			even = append(even, numbers[i])
		} else {
			odd = append(odd, numbers[i])
		}
	}

	return even, odd
}

// func SeparateEvenOdd(numbers []int) ([]int, []int) {
// 	even := []int{}
// 	odd := []int{}

// 	for _, num := range numbers {
// 		if num%2 == 0 {
// 			even = append(even, num)
// 		} else {
// 			odd = append(odd, num)
// 		}
// 	}

// 	return even, odd

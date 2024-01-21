package exercises

/*
Write a program that calculates and prints
the average of a slice of integers.
*/

func CalculateAverage(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

// func CalculateAverage(numbers []int) float64 {
//     sum := 0
//     for _, num := range numbers {
//         sum += num
//     }
//     return float64(sum) / float64(len(numbers))
// }

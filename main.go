package main

import "GO-exercises/exercises"

func main() {
	// //ex 1
	// numbers1 := []int{5, 10, 15, 20, 25}
	// average := exercises.CalculateAverage(numbers1)

	// fmt.Println("Average:", average)

	// //ex 2
	// numbers2 := []int{1, 2, 3, 2, 4, 5, 6, 4}
	// unique := exercises.UniqueElements(numbers2)

	// fmt.Println("Unique Elements:", unique)

	// //ex 3
	// inputStr := "Hello, World!"
	// reversed := exercises.ReverseString(inputStr)

	// fmt.Println("Reversed String:", reversed)

	// //ex 4
	// numbers3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// even, odd := exercises.SeparateEvenOdd(numbers3)

	// fmt.Println("Even Numbers:", even)
	// fmt.Println("Odd Numbers:", odd)

	// exercises.RoutineTest()

	nums := []int{2, 7, 11, 15}
	target := 9

	exercises.TwoSum(nums, target)
}

package main

import (
	kata "GO-exercises/codeWars"
	"GO-exercises/exercises"
	"fmt"
	"sort"
)

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

	// nums := []int{2, 7, 11, 15}
	// target := 9

	// exercises.TwoSum(nums, target)

	harry := exercises.Student{
		Name: "Harry",
		Age:  10,
	}

	numberList := []int{1, 2, 3, 4}

	numberList = append(numberList, 5)

	for _, num := range numberList {
		fmt.Println(num)
	}

	name := "john"
	name = "mary"
	fmt.Println(name)
	fmt.Println(numberList)

	ron := exercises.Student{
		Name: "Ron",
		Age:  10,
	}

	draco := exercises.Student{
		Name: "Draco",
		Age:  10,
	}

	harry.SetAge(18)
	ron.SetAge(19)
	draco.SetAge(69)

	studentList := []exercises.Student{harry, ron}
	fmt.Println(studentList)

	studentList = append(studentList, draco)
	fmt.Println(studentList)

	sort.Slice(studentList, func(a, b int) bool {
		return studentList[a].Age > studentList[b].Age
	})

	fmt.Println(studentList)

	s := "aaabbbbhaijjjm"
	kata.PrinterError(s)

	kata.Solve("abcd")
	kata.Solve("abcda")
	kata.Solve("abcdabc")
	kata.Solve("aaaa")
	kata.Solve("aa")
	kata.Solve("a")
}

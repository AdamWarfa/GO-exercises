package exercises

import "fmt"

func FromScratch() {
	//print
	fmt.Println("hello world")
	// variables
	const firstName string = "john"
	var lastName string = "smith"
	// var lastName = "john"
	// lastName := "john"
	fmt.Println(firstName + " " + lastName)

	//redclare
	lastName = "mary"
	fmt.Println(firstName + " " + lastName)

	//slices
	numberList := []int{1, 2, 3, 4}

	numberList = append(numberList, 5)

	doubleList := []int{}

	//JS - for (let i=0; i<numberList.length; i++) {

	for i := 0; i < len(numberList); i++ {
		doubleList = append(doubleList, numberList[i]*2)
	}
	fmt.Println(doubleList)

	newList := doubleLetters(numberList)
	fmt.Println(newList)
}

func doubleLetters(list []int) []int {
	doubleList := []int{}
	for i := 0; i < len(list); i++ {
		doubleList = append(doubleList, list[i]*2)
	}
	return doubleList
}

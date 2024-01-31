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

	// function creation

	//error handling
	errorList := []int{1, 2}

	newList, err := doubleLetters(numberList)
	wrongList, err := doubleLetters(errorList)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newList)
	fmt.Println(wrongList)
}

func doubleLetters(list []int) ([]int, error) {

	// move slice and loop
	// if else
	//return multiple

	if len(list) < 3 {
		err := fmt.Errorf("list must be at least 3 items long")
		return []int{}, err
	}

	doubleList := []int{}
	for i := 0; i < len(list); i++ {
		if i < 3 {
			doubleList = append(doubleList, list[i]*2)
		} else {
			doubleList = append(doubleList, list[i])
		}
	}
	return doubleList, nil
}

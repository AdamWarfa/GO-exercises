package exercises

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	House string
	Color string
}

func (student *Student) SetAge(age int) {
	student.Age = age
}

func (student *Student) SetColor() string {
	switch student.House {
	case "Gry":
		return "Red"
	case "Huf":
		return "Yellow"
	case "Rav":
		return "Blue"
	case "Sly":
		return "Green"
	default:
		return "Gray"
	}
}

func TestStudents() {
	harry := Student{
		Name:  "Harry",
		Age:   10,
		House: "Gry",
		Color: "Black",
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

	ron := Student{
		Name:  "Ron",
		Age:   10,
		House: "Gry",
	}

	draco := Student{
		Name:  "Draco",
		Age:   10,
		House: "Sly",
	}

	harry.SetAge(18)
	ron.SetAge(19)
	draco.SetAge(69)
	studentList := []Student{harry, ron}
	fmt.Println(studentList)

	studentList = append(studentList, draco)
	fmt.Println(studentList)

	sort.Slice(studentList, func(a, b int) bool {
		return studentList[a].Age > studentList[b].Age
	})

	// for _, student := range studentList {
	// 	student.Color = student.SetColor()
	// }

	for i := 0; i < len(studentList); i++ {
		studentList[i].Color = studentList[i].SetColor()
	}

	fmt.Println(studentList)
}

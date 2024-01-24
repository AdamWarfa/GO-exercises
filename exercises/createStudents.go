package exercises

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

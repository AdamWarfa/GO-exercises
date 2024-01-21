package exercises

type Student struct {
	Name string
	Age  int
}

func (student *Student) SetAge(age int) {
	student.Age = age
}

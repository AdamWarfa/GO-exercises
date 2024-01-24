package exercises

import (
	"fmt"
	"slices"
)

type MinStack struct {
	Stack []int
}

func Constructor() MinStack {
	this := MinStack{
		Stack: []int{},
	}
	return this
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
}

func (this *MinStack) Pop() {
	this.Stack = slices.Delete(this.Stack, len(this.Stack)-1, len(this.Stack))

}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return slices.Min(this.Stack)
}

func (this *MinStack) TestMinStack() {
	this.Push(4)
	fmt.Println(this)
	this.Push(-9)
	fmt.Println(this)
	this.Push(5)
	fmt.Println(this)
	fmt.Println(this.Top())
	fmt.Println(this.GetMin())
	this.Pop()
	fmt.Println(this)

}

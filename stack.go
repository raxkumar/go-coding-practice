package main

import "fmt"

type Stack[T any] struct {
	Top       T
	StackList []T
}

func (s *Stack[T]) Push(i T) {
	s.Top = i
	s.StackList = append(s.StackList, i)

}

func (s *Stack[T]) TopValue() {
	fmt.Println(s.Top)
}

func (s *Stack[T]) Len() {
	fmt.Println(len(s.StackList))
}

func (s *Stack[T]) IsEmpty() {
	if len(s.StackList) == 0 {
		fmt.Println("EMPTY")
	} else {
		fmt.Println("NOT EMPTY")
	}
}

func (s *Stack[T]) Pop() {
	temp := s.StackList[len(s.StackList)-1]

	s.StackList = s.StackList[:len(s.StackList)-1]
	s.Top = s.StackList[len(s.StackList)-1]
	fmt.Println("Value popped from stack list:", temp)
}

func (s *Stack[T]) display() {
	fmt.Println("STACK")
	fmt.Println("")
	fmt.Println(" _____ ")
	for i := len(s.StackList) - 1; i >= 0; i-- {
		fmt.Println("| ", s.StackList[i], " |")
		fmt.Println(" _____ ")
	}

}

func StackFunc() {

	stack := Stack[int]{}

	// stack.Push(1)
	// stack.Push(2)
	// stack.Push(3)
	// stack.Push(4)

	stack.TopValue()

	stack.Len()

	stack.IsEmpty()
	stack.Pop()
	stack.TopValue()
	stack.display()

	stack1 := Stack[string]{}

	stack1.Push("1")
	stack1.Push("2")
	stack1.Push("3")
	stack1.Push("4")

	stack1.TopValue()

	stack1.Len()

	stack1.IsEmpty()
	stack1.Pop()
	stack1.TopValue()
	stack1.display()

}

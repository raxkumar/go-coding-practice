package main

import "fmt"

type DNode struct {
	value int
	left  *DNode
	right *DNode
}

type DoubleLinkedList struct {
	Head *DNode
	Tail *DNode
}

func (l *DoubleLinkedList) push(value int) {

	newNode := &DNode{value, nil, nil}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	temp := l.Head
	for temp.right != nil {
		temp = temp.right
	}
	// 1-->2-->3
	temp.right = newNode

	// 1<--2<--3
	temp.right.left = temp

	// assign last node to tail
	l.Tail = temp.right

}

func (sll *DoubleLinkedList) display() {
	temp := sll.Head
	for {
		fmt.Print(temp.value)
		if temp.right == nil {
			break
		}
		fmt.Print(`-->`)
		temp = temp.right
	}
	fmt.Println()
}

func (sll *DoubleLinkedList) displayR() {
	temp := sll.Tail
	for {
		fmt.Print(temp.value)
		if temp.left == nil {
			break
		}
		fmt.Print(`-->`)
		temp = temp.left
	}
	fmt.Println()
}

func DoubleLinkedListFunc() {
	dll := &DoubleLinkedList{}

	dll.push(1)
	dll.push(2)
	dll.push(3)
	dll.push(3)
	dll.push(3)
	dll.push(5)

	dll.display()
	dll.displayR()
}

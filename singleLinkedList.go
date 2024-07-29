package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type SingleLinkedList struct {
	Head *Node
}

func (sll *SingleLinkedList) insert(value int) {
	newNode := &Node{Value: value, Next: nil}

	if sll.Head == nil {
		sll.Head = newNode
		return
	}
	temp := sll.Head

	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newNode
}

func (sll *SingleLinkedList) display() {
	temp := sll.Head
	for {
		fmt.Print(temp.Value)
		if temp.Next == nil {
			break

		}
		fmt.Print(`-->`)
		temp = temp.Next
	}
	fmt.Println()
}

func (sll *SingleLinkedList) search(value int) {
	if sll.Head.Value == value {
		fmt.Print(value, " FOUND\n")
		return
	}

	temp := sll.Head.Next
	for {
		if temp.Value == value {
			fmt.Print(value, " FOUND\n")
		}
		if temp.Next == nil {
			break
		}
		temp = temp.Next

	}

}

func (sll *SingleLinkedList) delete(value int) {
	if sll.Head.Value == value {
		sll.Head = sll.Head.Next
		return
	}
	temp := sll.Head.Next
	for {
		if temp.Value == value {
			sll.Head.Next = temp.Next
			return
		}
		if temp.Next.Value == value {
			temp.Next = temp.Next.Next
			return
		}
		if temp.Next == nil {
			break
		}
		temp = temp.Next

	}

}

func SingleLinkedListFunc() {

	ssl := SingleLinkedList{}

	ssl.insert(1)
	ssl.insert(2)
	ssl.insert(3)
	ssl.insert(4)
	ssl.insert(10)

	ssl.display()

	ssl.search(2)
	ssl.search(11)
	ssl.search(1)
	ssl.search(3)

	ssl.delete(10)

	ssl.display()

}

package main

import (
	"fmt"
)

type Node struct {
	value interface {}
	next *Node
}

type List struct {
	head *Node
}

func (l *List) add(v interface{}) {
	node := &Node { value: v, next: nil }

	if l.head == nil {
		l.head = node
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func display(l *List) {
	if l.head == nil {
		fmt.Println("Empty list")
	}

	current := l.head
	for current != nil {
		fmt.Println("The value of the node is: ", current.value)
		current = current.next
	}
}

func main() {
	list := List{}
	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)
	display(&list)
}
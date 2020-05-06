package main

import (
	"fmt"
)

// Node ...
type Node struct {
	data int
	next *Node
}

type linkedList struct {
	len  int
	head *Node
}

func initList() *linkedList {
	return &linkedList{}
}

func main() {
	head := initList()
	for i := 10; i < 100; i = i + 10 {
		head.addFront(i)
		head.addBack(i + 10)
	}
	head.display()
	head.reverseLinkedList()
	head.display()
}

func (list *linkedList) reverseLinkedList() {
	current := list.head
	if current == nil {
		return
	}
	var prev *Node
	var next *Node
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev
}

func (list *linkedList) addFront(val int) {
	// Getting Instance of New Node
	node := &Node{
		data: val,
	}

	// Checking if List is NULL
	if list.head == nil {
		list.head = node
	} else {
		node.next = list.head
		list.head = node
	}
	list.len++
	return
}

func (list *linkedList) addBack(val int) {
	node := &Node{
		data: val,
	}
	// Checking if List is NULL
	if list.head == nil {
		list.head = node
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	list.len++
	return
}

func (list *linkedList) display() {
	current := list.head
	for current != nil {
		fmt.Printf("%v\t", current.data)
		current = current.next
	}
	fmt.Println()
}

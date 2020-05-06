package main

import "fmt"

// Node ... Declaring Node Structure
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Tree ... Declaring Tree Structure
type Tree struct {
	root *Node
}

func (root *Node) preOrder() {
	if root == nil {
		return
	}
	fmt.Print(root.data, " ")
	root.left.preOrder()
	root.right.preOrder()
}

func main() {

	rlroot := &Node{
		data: 60,
	}
	rrroot := &Node{
		data: 70,
	}

	llroot := &Node{
		data: 40,
	}
	lrroot := &Node{
		data: 50,
	}
	rroot := &Node{
		data:  30,
		left:  rlroot,
		right: rrroot,
	}

	lroot := &Node{
		data:  20,
		left:  llroot,
		right: lrroot,
	}

	root := &Node{
		data:  10,
		left:  lroot,
		right: rroot,
	}

	root.preOrder()
	fmt.Println()

}

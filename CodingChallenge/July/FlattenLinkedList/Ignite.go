package main

// NodeX ...
type NodeX struct {
	val   int
	Next  *NodeX
	Prev  *NodeX
	Child *NodeX
}

func flatten(root *NodeX) *NodeX {
	Stack := make([]*NodeX, 0)
	makeFlatten(root, Stack)
	return root
}

func makeFlatten(head *NodeX, stack []*NodeX) {
	if head == nil {
		return
	}
	if head.Next == nil {
		// Checking if Last Node has Node
		if head.Child != nil {
			head.Next = head.Child
			head.Child.Prev = head
			head.Child = nil
			makeFlatten(head.Next, stack)
		}
		if len(stack) != 0 {
			found := stack[len(stack)-1]
			head.Next = found
			found.Prev = head
			stack = stack[:len(stack)-1]
			makeFlatten(head.Next, stack)
		}
		return
	}

	if head.Child != nil {
		stack = append(stack, head.Next)
		head.Next = head.Child
		head.Child.Prev = head
		head.Child = nil
	}
	makeFlatten(head.Next, stack)
}

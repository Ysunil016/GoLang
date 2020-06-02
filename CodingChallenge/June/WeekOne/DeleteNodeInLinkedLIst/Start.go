package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{10, &ListNode{20, &ListNode{30, &ListNode{}}}}
	deleteNode(list)
}

func deleteNode(node *ListNode) {
	PrevNode := &ListNode{}
	for node.Next != nil {
		node.Val = node.Next.Val
		PrevNode = node
		node = node.Next
	}
	PrevNode.Next = nil
}

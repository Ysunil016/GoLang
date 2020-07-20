package main

func main() {

}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	result := head
	if head == nil {
		return nil
	}

	if head.Val == val {
		result = head.Next
		head = head.Next
	}
	var PrevNode *ListNode
	for head != nil {
		if head.Val == val {
			if PrevNode == nil {
				result = head.Next
			} else {
				PrevNode.Next = head.Next
			}
		} else {
			PrevNode = head
		}
		head = head.Next
	}
	return result
}

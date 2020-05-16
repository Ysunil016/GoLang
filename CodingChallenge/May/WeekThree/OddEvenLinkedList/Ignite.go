package main

import "fmt"

// ListNode ...
type ListNode struct {
	data int
	next *ListNode
}

func main() {

	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	// head.displayList()
	head = oddEvenList(head)
	head.displayList()
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	oddMain := &ListNode{-1, nil}
	evenMain := &ListNode{-1, nil}
	oddCurrent := oddMain
	evenCurrent := evenMain
	for head != nil {
		oddCurrent.next = head
		oddCurrent = oddCurrent.next
		if head.next != nil {
			evenCurrent.next = head.next
			evenCurrent = evenCurrent.next
		} else {
			evenCurrent.next = nil
			break
		}
		head = head.next.next
	}
	oddCurrent.next = evenMain.next
	return oddMain.next
}

func oddEvenListOptamized(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	evenHead := odd.next
	even := evenHead

	for even != nil && even.next != nil {
		odd.next = even.next
		odd = odd.next
		even.next = odd.next
		even = even.next
	}
	odd.next = evenHead
	return head

}

func (head *ListNode) displayList() {
	for head != nil {
		fmt.Print(head.data, "->")
		head = head.next
	}
	fmt.Println()
}

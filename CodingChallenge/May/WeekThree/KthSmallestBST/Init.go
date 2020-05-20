package main

import (
	"fmt"
)

// Count ..
var Count int = 0

// KthSmallest ..
var KthSmallest int = -1

func main() {
	_1 := TreeNode{1, nil, nil}
	_2 := TreeNode{2, &_1, nil}
	_4 := TreeNode{4, nil, nil}
	_3 := TreeNode{3, &_2, &_4}
	_6 := TreeNode{6, nil, nil}
	root := TreeNode{5, &_3, &_6}
	kthSmallest(&root, 3)
	fmt.Println("Value ", KthSmallest)
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	Count = 0
	KthSmallest = -1
	preOrder(root, k)
	return KthSmallest
}

func preOrder(root *TreeNode, K int) {
	if root == nil {
		return
	}
	preOrder(root.Left, K)
	Count++
	if Count == K {
		KthSmallest = root.Val
	}
	preOrder(root.Right, K)
}

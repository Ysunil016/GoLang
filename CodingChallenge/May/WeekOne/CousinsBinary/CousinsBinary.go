package main

import "fmt"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root4 := TreeNode{4, nil, nil}
	root5 := TreeNode{5, nil, nil}
	root3 := TreeNode{3, nil, &root5}
	root2 := TreeNode{2, nil, &root4}
	root := TreeNode{1, &root2, &root3}
	fmt.Println(isCousins(&root, 4, 5))
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	type CousinS struct {
		node   *TreeNode
		parent *TreeNode
	}

	queue := []*TreeNode{root}
	for {
		size := len(queue)
		if size == 0 {
			break
		}

		X := CousinS{}
		Y := CousinS{}

		for size > 0 {
			element := queue[0]
			queue = queue[1:]
			if element.Left != nil {
				queue = append(queue, element.Left)
				if element.Left.Val == x {
					X = CousinS{element.Left, element}
				}
				if element.Left.Val == y {
					Y = CousinS{element.Left, element}
				}
			}

			if element.Right != nil {
				queue = append(queue, element.Right)
				if element.Right.Val == x {
					X = CousinS{element.Right, element}
				}
				if element.Right.Val == y {
					Y = CousinS{element.Right, element}
				}
			}

			size--
		}
		if X.node != nil && Y.node != nil {
			if X.parent != Y.parent {
				return true
			}
		}

	}
	return false
}

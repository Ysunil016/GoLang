package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	Arr := []int{4, 1, 2, 3, 6}
	bstFromPreorder(Arr)
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{preorder[0], nil, nil}
	for i := 1; i < len(preorder); i++ {
		makeBST(root, preorder[i])
	}
	return root
}
func makeBST(root *TreeNode, Val int) {
	if root == nil {
		return
	}
	if Val <= root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val, nil, nil}
		} else {
			makeBST(root.Left, Val)
		}
	}
	if Val > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{Val, nil, nil}
		} else {
			makeBST(root.Right, Val)
		}
	}
}

package main

func main() {

}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root.Val == val {
		return root
	}
	return hasFound(root, val)
}

func hasFound(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return hasFound(root.Left, val)
	} else {
		return hasFound(root.Right, val)
	}
}

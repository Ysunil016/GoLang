package main

// TreeNode ...
type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

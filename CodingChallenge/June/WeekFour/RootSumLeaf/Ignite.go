package main

func main() {

}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Result ...
var Result int

func sumNumbers(root *TreeNode) int {
	Result = 0
	traverse(root, 0)
	return Result
}

func traverse(root *TreeNode, Sum int) {
	if root == nil {
		return
	}
	Sum *= 10
	Sum += root.Val
	if root.Left == nil && root.Right == nil {
		Result += Sum
		return
	}
	traverse(root.Left, Sum)
	traverse(root.Right, Sum)
}

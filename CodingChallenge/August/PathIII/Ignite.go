package main

func main() {

}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var finalCout int = 0

func pathSum(root *TreeNode, sum int) int {
	finalCout = 0
	forEach(root, sum)
	return finalCout
}

func forEach(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	findPath(root, sum, root.Val)

	forEach(root.Left, sum)
	forEach(root.Right, sum)
}

func findPath(root *TreeNode, finalSum int, curSum int) {
	if root == nil {
		return
	}
	if finalSum == curSum {
		finalCout++
	}

	if root.Left != nil {
		findPath(root.Left, finalSum, curSum+root.Left.Val)
	}
	if root.Right != nil {
		findPath(root.Right, finalSum, curSum+root.Right.Val)
	}

}

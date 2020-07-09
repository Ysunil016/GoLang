package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

var hash map[int]int
var maxWidth int

func widthOfBinaryTree(root *TreeNode) int {
	hash = map[int]int{}
	maxWidth = 0
	getWidth(root, 0, 0)
	return maxWidth
}

func getWidth(root *TreeNode, depth, position int) {
	if root == nil {
		return
	}
	_, is := hash[depth]
	if !is {
		hash[depth] = position
	}
	maxWidth = getMax(maxWidth, position-hash[depth]+1)

	getWidth(root.Left, depth+1, position*2)
	getWidth(root.Right, depth+1, position*2+1)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

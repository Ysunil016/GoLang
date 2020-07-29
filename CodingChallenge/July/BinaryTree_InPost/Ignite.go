package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// Index ...
type Index int

func buildTree(inorder []int, postorder []int) *TreeNode {
	pI := len(inorder) - 1
	return buildUtils(inorder, postorder, 0, len(inorder)-1, pI)
}

func buildUtils(in []int, post []int, inStart int, inEnd int, pIndex int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	tNode := &TreeNode{Val: post[pIndex]}
	pIndex--

	if inStart == inEnd {
		return tNode
	}

	iIndex := searchIndex(in, tNode.Val)

	tNode.Right = buildUtils(in, post, inStart, iIndex-1, pIndex)
	tNode.Left = buildUtils(in, post, iIndex+1, inEnd, pIndex)
	return tNode
}

func searchIndex(ar []int, val int) int {
	for i, V := range ar {
		if V == val {
			return i
		}
	}
	return -1
}

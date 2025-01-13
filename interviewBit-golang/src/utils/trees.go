package utils

type TreeNode struct {
	Left  *TreeNode
	Value int
	Right *TreeNode
}

func TreeNode_new(val int) *TreeNode {
	node := new(TreeNode)
	node.Left = nil
	node.Right = nil
	node.Value = val
	return node
}

// func GenerateBinaryTree(A []int) *TreeNode {

// 	if len(A) == 0 {
// 		return nil
// 	}

// 	root := TreeNode_new(A[0])

// 	return root
// }

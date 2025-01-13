package trees

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/inorder-traversal/
*/

// T(n) : O(n) , S(n) : O(n)
func InorderTraversalRecursive(A *utils.TreeNode) []int {
	res := make([]int, 0)
	solveInOrderRecursive(A, &res)
	return res
}

func solveInOrderRecursive(A *utils.TreeNode, res *[]int) {
	if A == nil {
		return
	}
	solveInOrderRecursive(A.Left, res)
	*res = append(*res, A.Value)
	solveInOrderRecursive(A.Right, res)
}

// T(n) : O(n) , S(n) : O(n)
func InorderTraversalIterativeV1(A *utils.TreeNode) []int {

	res := make([]int, 0)
	stack := make([](*utils.TreeNode), 0)
	node := A
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Value)

		node = node.Right
	}

	return res
}

// T(n) : O(n) , S(n) : O(n)
func InorderTraversalIterativeV2(A *utils.TreeNode) []int {

	res := make([]int, 0)
	stack := make([](*utils.TreeNode), 0)
	node := A
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Right
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Value)

		node = node.Left
	}

	n := len(res)
	for i := 0; i < n/2; i++ {
		t := res[i]
		res[i] = res[n-i-1]
		res[n-i-1] = t
	}

	return res
}

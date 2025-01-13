package trees

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/postorder-traversal/
*/

// T(n) : O(n) , S(n) : O(n)
func PostOrderTraversalRecursive(A *utils.TreeNode) []int {
	res := make([]int, 0)
	solvePostOrderRecursive(A, &res)
	return res
}

func solvePostOrderRecursive(A *utils.TreeNode, res *[]int) {
	if A == nil {
		return
	}
	solvePostOrderRecursive(A.Left, res)
	solvePostOrderRecursive(A.Right, res)
	*res = append(*res, A.Value)
}

// T(n) : O(n) , S(n) : O(n)
func PostorderTraversalIterativeV1(A *utils.TreeNode) []int {

	res := make([]int, 0)
	stack := make([](*utils.TreeNode), 0)
	node := A
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			res = append(res, node.Value)
			node = node.Right
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
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

// T(n) : O(n) , S(n) : O(n)
func PostorderTraversalIterativeV2(A *utils.TreeNode) []int {

	res := make([]int, 0)
	stack := make([](*utils.TreeNode), 0)
	stack = append(stack, A)
	var prev *utils.TreeNode = nil
	for len(stack) > 0 {
		A = stack[len(stack)-1]
		if prev == nil || prev.Left == A || prev.Right == A {
			if A.Left != nil {
				stack = append(stack, A.Left)
			} else if A.Right != nil {
				stack = append(stack, A.Right)
			} else {
				stack = stack[:len(stack)-1]
				res = append(res, A.Value)
			}
		} else if A.Left == prev {
			if A.Right != nil {
				stack = append(stack, A.Right)
			} else {
				stack = stack[:len(stack)-1]
				res = append(res, A.Value)
			}
		} else if A.Right == prev {
			stack = stack[:len(stack)-1]
			res = append(res, A.Value)
		}
		prev = A
	}

	return res
}

package trees

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/preorder-traversal/
*/

// T(n) : O(n) , S(n) : O(n)
func PreorderTraversalRecursive(A *utils.TreeNode) []int {
	if A == nil {
		return []int{}
	}

	res := make([]int, 0)
	solvePreOrderRecursive(A, &res)

	return res
}

func solvePreOrderRecursive(A *utils.TreeNode, res *[]int) {
	if A == nil {
		return
	}

	*res = append(*res, A.Value)
	solvePreOrderRecursive(A.Left, res)
	solvePreOrderRecursive(A.Right, res)
}

// T(n) : O(n) , S(n) : O(n)
func PreOrderTraversalIterativeV1(A *utils.TreeNode) []int {
	if A == nil {
		return []int{}
	}

	res := make([]int, 0)
	stack := make([](*utils.TreeNode), 0)
	node := A
	for len(stack) > 0 || node != nil {
		for node != nil {
			res = append(res, node.Value)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node = node.Right
	}

	return res
}

// T(n) : O(n) , S(n) : O(n)
func PreOrderTraversalIterativeV2(A *utils.TreeNode) []int {

	if A == nil {
		return []int{}
	}

	res := make([]int, 0)
	stack := make([](*utils.TreeNode), 0)
	stack = append(stack, A)
	var prev *utils.TreeNode = nil
	for len(stack) > 0 {
		A = stack[len(stack)-1]
		if prev == nil || prev.Left == A || prev.Right == A {
			if A.Right != nil {
				stack = append(stack, A.Right)
			} else if A.Left != nil {
				stack = append(stack, A.Left)
			} else {
				stack = stack[:len(stack)-1]
				res = append(res, A.Value)
			}
		} else if A.Right == prev {
			if A.Left != nil {
				stack = append(stack, A.Left)
			} else {
				stack = stack[:len(stack)-1]
				res = append(res, A.Value)
			}
		} else if A.Left == prev {
			stack = stack[:len(stack)-1]
			res = append(res, A.Value)
		}
		prev = A
	}

	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-i-1] = res[n-i-1], res[i]
	}

	return res
}

package stacksAndQueues

import "interviewBit/src/utils"

/*
	Problem : https://www.interviewbit.com/problems/rain-water-trapped/

	Approach 0 : Brute Force
	Approach 1 : Using Extra Memory to calculate the rightMax array
	Approach 2 : Using Stack Data Structure
	Approach 3 : Using Two Pointers Approach
	Approach 4 : Using Divide & Conquer.
*/

// reference --> https://www.enjoyalgorithms.com/blog/trapping-rain-water

// T(n) : O(n2), S(n) : O(1)
func RainWaterTrapped(A []int) int {

	n, trappedWater := len(A), 0
	for i := 0; i < n; i++ {
		leftMax := 0
		for j := i; j >= 0; j-- {
			if leftMax < A[j] {
				leftMax = A[j]
			}
		}

		rightMax := 0
		for j := i; j < n; j++ {
			if rightMax < A[j] {
				rightMax = A[j]
			}
		}

		trappedWater += utils.MinOfIntsOrFloats[int](leftMax, rightMax) - A[i]
	}

	return trappedWater
}

// T(n) : O(n), S(n) : O(n)
func RainWaterTrapped1(A []int) int {

	n, trappedWater, leftMax := len(A), 0, 0
	rightMax := make([]int, n)
	rightMax[n-1] = A[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = utils.MaxOfIntsOrFloats[int](rightMax[i+1], A[i])
	}

	for i := 0; i < n; i++ {
		leftMax = utils.MaxOfIntsOrFloats[int](leftMax, A[i])
		trappedWater += utils.MinOfIntsOrFloats[int](leftMax, rightMax[i]) - A[i]
	}

	return trappedWater
}

func RainWaterTrapped2(A []int) int {

	n, stack, i := len(A), make([]int, 0), 0
	trappedWater := 0
	for i < n {
		if len(stack) == 0 || A[i] < A[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// if len(stack) > 0 {
			trappedWater -= (A[top] - A[i])
			// }
		}
	}

	return trappedWater
}

func RainWaterTrapped3(A []int) int {
	trappedWater := 0
	return trappedWater
}

func RainWaterTrapped4(A []int) int {
	trappedWater := 0
	return trappedWater
}

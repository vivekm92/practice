package stacksAndQueues

import "math"

/*
	Problem :
*/

// T(n) : O(n2), S(n) : O(1)
func LargestRectangleInHistogram(A []int) int {

	n := len(A)
	currArea, maxArea := math.MinInt32, math.MinInt32
	for i := 0; i < n; i++ {
		l := i
		for l >= 0 && A[l] >= A[i] {
			l--
		}

		r := i
		for r < n && A[r] >= A[i] {
			r++
		}

		currArea = (r - l - 1) * A[i]
		if maxArea < currArea {
			maxArea = currArea
		}
	}

	return maxArea
}

// T(n) : O(n), S(n) : O(n)
func LargestRectangleInHistogram1(A []int) int {

	n, stack := len(A), make([]int, 0)
	currArea, maxArea, i := 0, 0, 0
	for i < n {
		if len(stack) == 0 || A[i] >= A[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				currArea = i * A[top]
			} else {
				currArea = A[top] * (i - stack[len(stack)-1] - 1)
			}

			if maxArea < currArea {
				maxArea = currArea
			}
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			currArea = i * A[top]
		} else {
			currArea = A[top] * (i - stack[len(stack)-1] - 1)
		}

		if maxArea < currArea {
			maxArea = currArea
		}
	}

	return maxArea
}

package arrayProblems

import "sort"

/*
	Problem : https://www.interviewbit.com/problems/sort-array-with-squares/
	LeetCode : https://leetcode.com/problems/squares-of-a-sorted-array/description/
*/

// SortArrayWithSquares sorts the input array A of n integers by the squares of its elements in ascending order
// Time complexity: O(n), Space complexity: O(n)
//

func SortArrayWithSquares(A []int) []int {
	res := make([]int, 0, len(A)) // preallocate the space for the result

	i, j := 0, len(A)-1 // two pointers, one from the beginning and one from the end of A
	for i <= j {        // loop until we reach the middle
		if A[i]*A[i] > A[j]*A[j] { // compare the squares and append the smaller one
			res = append(res, A[i]*A[i])
			i++
		} else {
			res = append(res, A[j]*A[j])
			j--
		}
	}

	for i := 0; i < len(A)/2; i++ { // reverse the result
		res[i], res[len(A)-1-i] = res[len(A)-1-i], res[i]
	}

	return res // return the sorted array
}

// SortArrayWithSquares1 sorts the input array A of n integers by the squares of its elements in ascending order
// Time complexity: O(n), Space complexity: O(n)
//
// The algorithm works by finding the first negative element in the array A using binary search and then
// iterating through the array from the left and right ends simultaneously, comparing the squares of the elements
// at these indices and appending the smaller square to the result. The result is then returned reversed.
func SortArrayWithSquares1(A []int) []int {
	// find the index of the first negative element in A using binary search
	n := len(A)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if A[mid] < 0 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	// l is now the index of the first negative element in A, or n if there is no negative element
	l = l % n

	// two pointers, one from the left and one from the right end of A
	l1, l2, i := l-1, l, 0
	res := make([]int, n)
	for l1 >= 0 || l2 < n {
		if l1 >= 0 && l2 < n { // if both indices are valid
			if A[l1]*A[l1] < A[l2]*A[l2] { // compare the squares of the elements at these indices
				res[i] = A[l1] * A[l1] // append the smaller square to the result
				l1 -= 1                // move the left pointer left
			} else {
				res[i] = A[l2] * A[l2] // append the smaller square to the result
				l2 += 1                // move the right pointer right
			}
		} else if l1 >= 0 { // if the left index is valid but the right index is not
			res[i] = A[l1] * A[l1] // append the square at the left index to the result
			l1 -= 1                // move the left pointer left
		} else { // if the right index is valid but the left index is not
			res[i] = A[l2] * A[l2] // append the square at the right index to the result
			l2 += 1                // move the right pointer right
		}
		i += 1 // increment the result index
	}

	return res
}

// T(n) : O(nlogn), S(n) : O(n)
func SortArrayWithSquares2(A []int) []int {

	n := len(A)
	for i := 0; i < n; i++ {
		A[i] = A[i] * A[i]
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	return A
}

/*
Approach 0 :

> Since array is in sorted order, the maximum value of square can be present at either end of the array only.
> We can append the square values in non increasing order, and later reverse the array to arrive at solution.

Approach 1 :

> Find the idx, such that A[idx] >= 0 ... since the array is sorted we can use binary serach for this.
> Now we can track and compare the square values and push the same in result array.

Approach 2 :



*/

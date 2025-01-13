package greedy

import (
	"sort"
)

/*
	Problem : https://www.interviewbit.com/problems/distribute-candy/
*/

// T(n) : O(nlogn) ; S(n) : O(n)
func DistributeCandy(A []int) int {

	n := len(A)
	if n <= 1 {
		return n
	}

	arr := make([][]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, []int{A[i], i})
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] < arr[j][0]
	})

	res, temp := 0, make([]int, n)
	for i := 0; i < n; i++ {
		curr := arr[i]
		if curr[1] == 0 {
			if curr[0] <= A[curr[1]+1] {
				temp[curr[1]] = 1
			} else {
				temp[curr[1]] = temp[curr[1]+1] + 1
			}
		} else if curr[1] == n-1 {
			if curr[0] <= A[curr[1]-1] {
				temp[curr[1]] = 1
			} else {
				temp[curr[1]] = temp[curr[1]-1] + 1
			}
		} else {
			if curr[0] <= A[curr[1]+1] && curr[0] <= A[curr[1]-1] {
				temp[curr[1]] = 1
			} else if curr[0] > A[curr[1]+1] && curr[0] <= A[curr[1]-1] {
				temp[curr[1]] = temp[curr[1]+1] + 1
			} else if curr[0] <= A[curr[1]+1] && curr[0] > A[curr[1]-1] {
				temp[curr[1]] = temp[curr[1]-1] + 1
			} else {
				t := temp[curr[1]+1]
				if t < temp[curr[1]-1] {
					t = temp[curr[1]-1]
				}
				temp[curr[1]] = t + 1
			}
		}

		res += temp[curr[1]]
	}

	return res
}

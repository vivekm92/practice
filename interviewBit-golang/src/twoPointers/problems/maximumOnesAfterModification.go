package twoPointers

/*
	Problem : https://www.interviewbit.com/problems/maximum-ones-after-modification/
*/

// T(n) : O(n), S(n) : O(1)
func MaximumOnesAfterModification(A []int, B int) int {

	n, k := len(A), B
	pt1, pt2 := 0, 0
	for pt2 < n && k > 0 {
		if A[pt2] == 0 {
			k--
		}
		pt2++
	}

	res := pt2 - pt1
	for pt2 < n {

		if A[pt1] == 0 {
			pt1++
			k++
		} else {
			for pt1 < n && A[pt1] == 1 {
				pt1++
			}
		}

		for pt2 < n && k > 0 {
			if A[pt2] == 0 {
				k--
			}
			pt2++
		}

		for pt2 < n && A[pt2] == 1 {
			pt2++
		}

		if res < pt2-pt1 {
			res = pt2 - pt1
		}
	}

	return res
}

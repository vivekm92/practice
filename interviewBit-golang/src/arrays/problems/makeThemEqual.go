package arrayProblems

import "math"

/*
	Problem : https://www.interviewbit.com/problems/make-them-equal/
*/

// T(n) : O(nlogn) : S(n) : O(1)
func MakeThemEqual(A []int) int {

	m, minVal := make(map[int]int), math.MaxInt32
	for _, v := range A {
		m[v]++
		if minVal > v {
			minVal = v
		}
	}

	if len(m) == 1 {
		return 0
	}

	count := 0
	for len(m) > 1 {
		for k, v := range m {
			currCount, k1, step, flag := 0, k, 0, false
			for k1 > minVal {
				k1 /= 2
				step++
				flag = true
			}

			if flag {
				if _k, _ok := m[k1]; _ok {
					m[k1] = _k + v
				} else {
					m[k1] = v
				}
				if minVal > k1 {
					minVal = k1
				}

				currCount += step * v
				delete(m, k)
				count += currCount
			}
		}
	}

	return count
}

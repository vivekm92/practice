package stringProblems

import (
	"fmt"
)

/*
	problem : https://www.interviewbit.com/problems/bulls-and-cows/
*/

// T(n) : O(n), S(n) : O(n)
func BullsAndCows(A string, B string) string {

	n, nBulls, l1, l2 := len(A), 0, make(map[byte]int), make(map[byte]int)
	for i := 0; i < n; i++ {
		if A[i] == B[i] {
			nBulls++
		} else {
			l1[A[i]]++
			l2[B[i]]++
		}
	}

	nCows := 0
	for k, v := range l1 {
		if val, ok := l2[k]; ok {
			if val < v {
				nCows += val
			} else {
				nCows += v
			}
		}
	}
	return fmt.Sprintf("%vA%vB", nBulls, nCows)
}

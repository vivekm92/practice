package hashing

import (
	"strconv"
)

/*
	Problem : https://www.interviewbit.com/problems/colorful-number
*/

func findProduct(A string) int {

	res := 1
	for i := 0; i < len(A); i++ {
		if _val, err := strconv.Atoi(string(A[i])); err == nil {
			res *= _val
		}
	}

	return res
}

// T(n) : O(logN) ; S(n) : O(logN)
func ColorfulNumber(A int) int {

	m := make(map[int]int, 0)
	B := strconv.Itoa(A)
	for i := 0; i < len(B); i++ {
		for j := i; j < len(B); j++ {
			t := B[i : j+1]
			prod := findProduct(t)
			m[prod]++
			if _val, ok := m[prod]; ok && _val > 1 {
				return 0
			}
		}
	}

	return 1
}

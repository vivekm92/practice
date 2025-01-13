package stringProblems

import "strings"

/*
	Problem : https://www.interviewbit.com/problems/compare-version-numbers/
*/

func trim(A string) string {

	// Trim Leading Zeros
	i := 0
	for i < len(A) && A[i] == '0' {
		i++
	}
	A = A[i:]

	// Trim Zeros present at end
	i = len(A) - 1
	for i >= 0 && A[i] == '0' {
		i--
	}
	A = A[:i+1]

	// Handle cases like 12.0
	if len(A) > 0 && A[len(A)-1] == '.' {
		A = A[:len(A)-1]
	}

	return A
}

func compareSubstring(A, B string) int {

	if len(A) > len(B) {
		return 1
	} else if len(B) > len(A) {
		return -1
	}

	i := 0
	for i < len(A) {
		if A[i] > B[i] {
			return 1
		} else if A[i] < B[i] {
			return -1
		}
		i++
	}

	return 0
}

// T(n) : O(n) , S(n) : O(1)
func CompareVersionNumbers(A string, B string) int {

	A = trim(A)
	B = trim(B)

	a := strings.Split(A, ".")
	b := strings.Split(B, ".")

	i, j := 0, 0
	for i < len(a) || j < len(b) {
		if i >= len(a) {
			return -1
		} else if j >= len(b) {
			return 1
		}

		t := compareSubstring(a[i], b[j])
		if t != 0 {
			return t
		}
		i++
		j++
	}

	return 0
}

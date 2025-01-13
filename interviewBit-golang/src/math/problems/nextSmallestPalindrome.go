package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/next-smallest-palindrome/
*/

import (
	"strconv"
	"strings"
)

// T(n) : O(n) ; S(n) : O(n)
func NextSmallestPalindrome(A string) string {

	// Case 1 : All Digits are 9
	n := len(A)
	if allDigitsNine(A) {
		return "1" + strings.Repeat("0", n-1) + "1"
	}

	mid := n / 2
	i, j := mid-1, mid
	if n%2 != 0 {
		j = mid + 1
	}
	for i >= 0 && j < n && A[i] == A[j] {
		i--
		j++
	}

	leftSmaller := false
	if i < 0 || A[i] < A[j] {
		leftSmaller = true
	}

	if n%2 != 0 {
		A = A[:mid] + string(A[mid]) + reverseString((A[:mid]))
	} else {
		A = A[:mid] + reverseString(A[:mid])
	}

	if leftSmaller {
		if n%2 != 0 {
			A = incrementDigit(A, mid)
			A = A[:mid] + string(A[mid]) + reverseString((A[:mid]))
		} else {
			A = incrementDigit(A, mid-1)
			A = A[:mid] + reverseString(A[:mid])
		}
	}

	return A
}

func allDigitsNine(A string) bool {
	for _, v := range A {
		if v != '9' {
			return false
		}
	}
	return true
}

func reverseString(A string) string {
	runes, n := []rune(A), len(A)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes)
}

func incrementDigit(A string, idx int) string {

	carry := 1
	for idx >= 0 {
		v, _ := strconv.Atoi(string(A[idx]))
		v = v + carry

		A = A[:idx] + strconv.Itoa(v%10) + A[idx+1:]
		carry = v / 10
		if carry == 0 {
			return A
		}
		idx--
	}

	return A
}

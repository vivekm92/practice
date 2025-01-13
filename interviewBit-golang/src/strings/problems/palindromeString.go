package stringProblems

/*
problem : https://www.interviewbit.com/problems/palindrome-string/
*/

// T(n) : O(1), S(n) : O(1)
func IsValidChar(a byte) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '1' && a <= '9') {
		return true
	}

	return false
}

// T(n) : O(1), S(n) : O(1)
func IsEqualChar(a, b byte) bool {

	if a >= 'A' && a <= 'Z' {
		a = a - 'A' + 'a'
	}

	if b >= 'A' && b <= 'Z' {
		b = b - 'A' + 'a'
	}

	if a >= '1' && a <= '9' {
		a = a - '1' + 'a'
	}

	if b >= '1' && b <= '9' {
		b = b - '1' + 'a'
	}

	return a == b
}

// T(n) : O(n), S(n) : O(1)
func IsPalindrome(A string) int {

	n := len(A)
	i, j := 0, n-1
	for i < j {
		for i < n && !IsValidChar(A[i]) {
			i++
		}
		for j >= 0 && !IsValidChar(A[j]) {
			j--
		}

		if !IsEqualChar(A[i], A[j]) {
			return 0
		} else {
			i++
			j--
		}
	}

	return 1
}

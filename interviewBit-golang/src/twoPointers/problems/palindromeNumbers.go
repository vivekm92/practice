package twoPointers

/*
problem : https://www.interviewbit.com/problems/palindrome-numbers/
*/

// T(n) : O(logn), S(n) : O(1)
func IsPalindrome(n int) bool {

	t, r := n, 0
	for n > 0 {
		r = r*10 + n%10
		n = n / 10
	}

	return r == t
}

// T(n) : O(n), S(n) : O(n)
func CountPalindromeNumbers(A int, B int, C int) int {

	var t []int
	for i := A; i <= B; i++ {
		if IsPalindrome(i) {
			t = append(t, i)
		}
	}

	n := len(t)
	if n <= 1 {
		return n
	}

	j, maxCount := 0, 0
	for i := 0; i < n; i++ {
		for j < n && t[j]-t[i] <= C {
			j++
		}
		currCount := j - i
		if maxCount < currCount {
			maxCount = currCount
		}
	}

	return maxCount
}

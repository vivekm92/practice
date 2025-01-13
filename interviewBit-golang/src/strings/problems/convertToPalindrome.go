package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/convert-to-palindrome/
*/

// T(n) : O(n) ; S(n) : O(1)
func ConvertToPalindrome(A string) int {

	n := len(A)
	i, j := isPalindrome(A)
	if i >= j {
		if n%2 == 1 {
			return 1
		} else {
			return 0
		}
	}

	v1 := A[:i] + A[i+1:]
	v2 := A[:j] + A[j+1:]
	x1, y1 := isPalindrome(v1)
	x2, y2 := isPalindrome(v2)
	if x1 >= y1 || x2 >= y2 {
		return 1
	}

	return 0
}

func isPalindrome(A string) (int, int) {
	i, j := 0, len(A)-1
	for i < j && A[i] == A[j] {
		i++
		j--
	}
	return i, j
}

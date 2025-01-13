package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/remove-consecutive-characters/
*/

// T(n) : O(n) , S(n) : O(1)
func RemoveConsecutiveCharacters(A string, B int) string {

	i, n := 0, len(A)
	for i < n {
		j, count := i+1, 1
		for j < n && A[i] == A[j] {
			j++
			count++
		}

		if count == B {
			A = A[0:i] + A[j:]
			n = len(A)
		} else {
			i = j
		}
	}

	return A
}

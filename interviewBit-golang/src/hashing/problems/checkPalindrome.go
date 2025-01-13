package hashing

/*
	Problem : https://www.interviewbit.com/problems/check-palindrome/
*/

// T(n) : O(n), S(n) : O(n)
func CheckPalindrome(A string) int {

	lookup := make(map[rune]int, 0)
	for _, v := range A {
		lookup[v]++
	}

	eventCount, oddCount := 0, 0
	for _, v := range lookup {
		if v%2 == 0 {
			eventCount++
		} else {
			oddCount++
		}
	}

	if oddCount < 2 {
		return 1
	}
	return 0
}

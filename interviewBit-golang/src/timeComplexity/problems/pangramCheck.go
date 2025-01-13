package timeComplexity

/*
	Problem : https://www.interviewbit.com/problems/pangram-check/
*/

// T(n) : O(n), S(n) : O(1)
func CheckIfPanagram(A []string) int {

	lookup := make([]rune, 26)
	for _, v1 := range A {
		for _, v2 := range v1 {
			lookup[v2-'a']++
		}
	}

	for _, v := range lookup {
		if v < 1 {
			return 0
		}
	}

	return 1
}

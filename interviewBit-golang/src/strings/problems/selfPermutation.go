package stringProblems

/*
	problem : https://www.interviewbit.com/problems/self-permutation/
*/

// T(n) : O(n), S(n) : O(n)
func PermuteStrings(A string, B string) int {

	l1 := make(map[rune]int)
	for _, v := range A {
		l1[v]++
	}

	l2 := make(map[rune]int)
	for _, v := range B {
		l2[v]++
	}

	if len(l1) != len(l2) {
		return 0
	}

	for k, v := range l1 {
		if val, ok := l2[k]; !ok || val != v {
			return 0
		}
	}

	return 1
}

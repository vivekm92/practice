package arrayProblems

// T(n): O(n), S(n) : O(n)
func FindPermutation(A string, B int) []int {

	l, r := 1, B
	permutation := make([]int, 0)
	for _, v := range A {
		if v == 'I' {
			permutation = append(permutation, l)
			l++
		} else {
			permutation = append(permutation, r)
			r--
		}
	}
	permutation = append(permutation, l)

	return permutation
}

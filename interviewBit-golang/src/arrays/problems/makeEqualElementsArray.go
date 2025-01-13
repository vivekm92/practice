package arrayProblems

// T(n) : O(n), S(n) : O(n)
func MakeEqualElements(A []int, B int) int {

	n, lookup := len(A), make(map[int]int, 0)
	for _, i := range A {
		lookup[i]++
		lookup[i-B]++
		lookup[i+B]++
	}

	for _, v := range lookup {
		if v == n {
			return 1
		}
	}
	return 0
}

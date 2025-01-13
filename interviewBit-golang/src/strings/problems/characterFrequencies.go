package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/character-frequencies/
*/

// T(n) : O(n), S(n) : O(n)
func CharacterFrequencies(A string) []int {

	res, lookup := make([]int, 0), make(map[rune]int)
	for _, v := range A {
		if x, ok := lookup[v]; ok {
			lookup[v] = x + 1
		} else {
			res = append(res, int(v))
			lookup[v] = x + 1
		}
	}

	n := len(res)
	for i := 0; i < n; i++ {
		res[i] = lookup[rune(res[i])]
	}

	return res
}

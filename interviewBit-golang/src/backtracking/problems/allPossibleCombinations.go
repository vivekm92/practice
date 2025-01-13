package backtracking

import "sort"

/*
	Problem : https://www.interviewbit.com/problems/all-possible-combinations/
*/

func solveAllPossibleCombinations(A []string, idx int) []string {
	if idx == len(A)-1 {
		t := make([]string, 0)
		for _, v := range A[idx] {
			t = append(t, string(v))
		}
		return t
	}

	t := solveAllPossibleCombinations(A, idx+1)
	r := make([]string, 0)
	for _, v1 := range t {
		for _, v2 := range A[idx] {
			r = append(r, string(v2)+v1)
		}
	}

	return r
}

// T(n) : O() ; S(n) : O()
func AllPossibleCombinations(A []string) []string {
	if len(A) == 0 {
		return []string{}
	}
	res := solveAllPossibleCombinations(A, 0)
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}

package arrayProblems

import (
	"sort"
)

func OccurenceOfEachNumnber(A []int) []int {

	m := make(map[int]int, 0)
	for _, v := range A {
		m[v]++
	}

	var lookup []int
	for k := range m {
		lookup = append(lookup, k)
	}

	sort.Slice(lookup, func(i, j int) bool {
		return lookup[i] < lookup[j]
	})

	var res []int
	for _, v := range lookup {
		res = append(res, m[v])
	}

	return res
}
